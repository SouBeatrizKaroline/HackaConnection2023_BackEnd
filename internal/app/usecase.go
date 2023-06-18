package app

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/julioc98/citi/internal/domain"
)

type ShippingRepository interface {
	Save(ctx context.Context, filename string, shipping *domain.Shipping) error
}

type BacenGateway interface {
	CobPUT(ctx context.Context, txID string, shippingDetail domain.ShippingDetail) error
}

type ReturnStorage interface {
	Save(ctx context.Context, filename string, r domain.Return) error
}

// ShippingUseCase is a struct that represents a use case for CNAB file (pt-BR Remessa)
type ShippingUseCase struct {
	// ShippingRepository domain.ShippingRepository
	repo    ShippingRepository
	bacen   BacenGateway
	storage ReturnStorage
}

// NewShippingUseCase returns a new ShippingUseCase
func NewShippingUseCase(repo ShippingRepository, bacen BacenGateway, storage ReturnStorage) *ShippingUseCase {
	return &ShippingUseCase{
		repo:    repo,
		bacen:   bacen,
		storage: storage,
	}
}

// MainFlow is a function that represents the main flow of a CNAB file (pt-BR Remessa)
func (u *ShippingUseCase) MainFlow(ctx context.Context, filename string, shipping *domain.Shipping) error {

	err := u.repo.Save(ctx, filename, shipping)
	if err != nil {
		return err
	}

	r, err := u.makeReturn(ctx, filename, *shipping)
	if err != nil {
		return err
	}

	err = u.storage.Save(ctx, filename, *r)
	if err != nil {
		return err
	}

	return nil
}

func (u *ShippingUseCase) makeReturn(ctx context.Context, filename string, shipping domain.Shipping) (*domain.Return, error) {
	returnHeader, err := u.shippingHeaderToReturnHeader(ctx, shipping.Header)
	if err != nil {
		return nil, err
	}

	var details []domain.ReturnDetail

	for i, detail := range shipping.Detail {
		tmpID := uuid.New().String()
		if detail.Identificador != "" {
			tmpID = detail.Identificador
		}

		err := u.bacen.CobPUT(ctx, tmpID, detail)
		if err != nil {
			return nil, err
		}

		returnDetail := domain.ReturnDetail{
			TipoRegistro:           1,
			Identificador:          tmpID,
			TipoPessoaRecebedor:    detail.TipoPessoaRecebedor,
			CNPJRecebedor:          detail.CNPJRecebedor,
			AgenciaRecebedor:       detail.AgenciaRecebedor,
			ContaRecebedor:         detail.ContaRecebedor,
			Tipo:                   "",
			ChavePix:               detail.ChavePix,
			TipoCobranca:           detail.TipoCobranca,
			CodigoMovimento:        1,
			TimestampExpiracao:     detail.TimestampExpiracao,
			DataVencimento:         detail.DataVencimento,
			ValidadeAposVencimento: 0,
			ValorOriginal:          detail.ValorOriginal,
			TipoPessoaDevedor:      detail.TipoPessoaDevedor,
			CNPJDevedor:            detail.CNPJDevedor,
			NomeDevedor:            detail.NomeDevedor,
			SolicitacaoPagador:     detail.SolicitacaoPagador,
			ExclusivoPSPRecebedor:  detail.ExclusivoPSPRecebedor,
			DataMovimento:          0,
			CodigosErro:            "",
			Revisao:                2,
			NumeroSequencial:       i + 1,
		}

		details = append(details, returnDetail)

	}

	returnTrailer := u.shippingTrailerToReturnTrailer(shipping)

	r := &domain.Return{
		Header:  returnHeader,
		Detail:  details,
		Trailer: returnTrailer,
	}

	return r, nil
}

// shippingHeaderToReturnHeader is a function that converts a ShippingHeader to a ReturnHeader
func (u *ShippingUseCase) shippingHeaderToReturnHeader(ctx context.Context, sh domain.ShippingHeader) (domain.ReturnHeader, error) {
	const (
		YYYYMMDD = "20060102"
	)

	now := time.Now()
	date := now.Format(YYYYMMDD)
	dateInt, err := strconv.Atoi(date)
	if err != nil {
		fmt.Println("Error during conversion")
	}

	returnHeader := domain.ReturnHeader{
		TipoRegistro:             0,
		CodigoRetorno:            2,
		LiteralRetorno:           "RETORNO",
		CodigoServico:            2,
		LiteralServico:           "PIX",
		ISPB:                     sh.ISPB,
		TipoPessoaRecebedor:      sh.TipoPessoa,
		CNPJRecebedor:            sh.CNPJ,
		AgenciaRecebedor:         sh.Agencia,
		ContaRecebedor:           sh.Conta,
		TipoContaRecebedor:       sh.TipoConta,
		ChavePix:                 sh.ChavePix,
		DataGeracao:              dateInt,
		CodigoConvenio:           sh.CodigoConvenio,
		ExclusivoPSP:             sh.ExclusivoPSP,
		NomeRecebedor:            sh.NomeRecebedor,
		CodigosErro:              "",
		Brancos:                  "",
		NumeroSequencial:         1,
		VersaoArquivo:            2,
		NumeroSequencialRegistro: 1,
	}

	return returnHeader, nil
}

func (u *ShippingUseCase) shippingTrailerToReturnTrailer(s domain.Shipping) domain.ReturnTrailer {
	returnTrailer := domain.ReturnTrailer{
		TipoRegistro:     9,
		CodigoRetorno:    2,
		CodigoServico:    2,
		ISPB:             s.Header.ISPB,
		Brancos:          s.Header.Brancos,
		ValorTotal:       s.Trailer.ValorTotal,
		QtdeDetalhes:     s.Trailer.QtdeRegistros,    // Assuming QtdeRegistros corresponds to QtdeDetalhes in ReturnTrailer
		NumeroSequencial: s.Trailer.NumeroSequencial, // Assuming NumeroSequencial is the same field name in both structs
	}

	return returnTrailer
}
