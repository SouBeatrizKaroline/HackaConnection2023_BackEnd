package app

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/julioc98/citi/internal/domain"
)

type ShippingRepository interface {
	Save(ctx context.Context, filename string, shipping *domain.Shipping) error
}

type BacenGateway interface {
	CobPUT(ctx context.Context, txID string, shippingDetail domain.ShippingDetail) error
}

// ShippingUseCase is a struct that represents a use case for CNAB file (pt-BR Remessa)
type ShippingUseCase struct {
	// ShippingRepository domain.ShippingRepository
	repo  ShippingRepository
	bacen BacenGateway
}

// NewShippingUseCase returns a new ShippingUseCase
func NewShippingUseCase(repo ShippingRepository, bacen BacenGateway) *ShippingUseCase {
	return &ShippingUseCase{
		repo:  repo,
		bacen: bacen,
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

	log.Println(r)

	return nil
}

func (u *ShippingUseCase) makeReturn(ctx context.Context, filename string, shipping domain.Shipping) (*domain.Return, error) {
	returnHeader, err := u.shippingHeaderToReturnHeader(ctx, shipping.Header)
	if err != nil {
		return nil, err
	}

	var details []domain.ReturnDetail

	for _, detail := range shipping.Detail {
		tmpID := uuid.New().String()
		if detail.Identificador != "" {
			tmpID = detail.Identificador
		}

		err := u.bacen.CobPUT(ctx, tmpID, detail)
		if err != nil {
			return nil, err
		}

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
	returnHeader := domain.ReturnHeader{
		TipoRegistro:             0,
		CodigoRetorno:            2,
		LiteralRetorno:           "RETORNO",
		CodigoServico:            02,
		LiteralServico:           "PIX",
		ISPB:                     sh.ISPB,
		TipoPessoaRecebedor:      sh.TipoPessoa,
		CNPJRecebedor:            sh.CNPJ,
		AgenciaRecebedor:         sh.Agencia,
		ContaRecebedor:           sh.Conta,
		TipoContaRecebedor:       sh.TipoConta,
		ChavePix:                 sh.ChavePix,
		DataGeracao:              sh.DataGeracao, // TODO: time.NOW() format AAAAMMDD
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
