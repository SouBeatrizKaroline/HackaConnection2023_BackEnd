package domain

import (
	"io"
	"log"

	"github.com/julioc98/gocnab"
)

// Shipping is a struct that represents a CNAB 750 file (pt-BR Remessa)
type Shipping struct {
	Header     ShippingHeader
	Detail     []ShippingDetail
	Trailer    ShippingTrailer
	Additional []ShippingDetailAdditional
	Charge     []ShippingDetailCharge
}
type ShippingHeader struct {
	TipoRegistro             int    `cnab:"0,1"`
	Operacao                 int    `cnab:"1,2"`
	LiteralRemessa           string `cnab:"2,9"`
	CodigoServico            int    `cnab:"9,11"`
	LiteralServico           string `cnab:"11,26"`
	ISPB                     string `cnab:"26,34"`
	TipoPessoa               int    `cnab:"34,36"`
	CNPJ                     int    `cnab:"36,50"`
	Agencia                  int    `cnab:"50,54"`
	Conta                    int    `cnab:"54,74"`
	TipoConta                string `cnab:"74,78"`
	ChavePix                 string `cnab:"78,155"`
	DataGeracao              int    `cnab:"155,163"`
	CodigoConvenio           string `cnab:"163,193"`
	ExclusivoPSP             string `cnab:"193,253"`
	NomeRecebedor            string `cnab:"253,353"`
	Brancos                  string `cnab:"353,731"`
	NumeroSequencial         int    `cnab:"731,741"`
	VersaoLayout             int    `cnab:"741,744"`
	NumeroSequencialRegistro int    `cnab:"744,750"`
}

type ShippingDetail struct {
	TipoRegistro          int     `cnab:"0,1"`
	Identificador         string  `cnab:"1,36"`
	TipoPessoaRecebedor   int     `cnab:"36,38"`
	CNPJRecebedor         int     `cnab:"38,52"`
	AgenciaRecebedor      int     `cnab:"52,56"`
	ContaRecebedor        int     `cnab:"56,76"`
	TipoContaRecebedor    string  `cnab:"76,80"`
	ChavePix              string  `cnab:"80,157"`
	TipoCobranca          string  `cnab:"157,158"`
	CodOcorrencia         int     `cnab:"158,160"`
	TimestampExpiracao    int     `cnab:"160,174"`
	DataVencimento        int     `cnab:"174,182"`
	ValidadeAposVenc      int     `cnab:"182,186"`
	ValorOriginal         float64 `cnab:"186,203"`
	TipoPessoaDevedor     int     `cnab:"203,205"`
	CNPJDevedor           int     `cnab:"205,219"`
	NomeDevedor           string  `cnab:"219,359"`
	SolicitacaoPagador    string  `cnab:"359,499"`
	ExclusivoPSPRecebedor string  `cnab:"499,559"`
	Brancos               string  `cnab:"559,744"`
	NumeroSequencial      int     `cnab:"744,750"`
}

type ShippingDetailAdditional struct {
	TipoRegistro     int    `cnab:"0,1"`
	Identificador    string `cnab:"1,36"`
	Nome             string `cnab:"36,86"`
	Valor            string `cnab:"86,286"`
	Nome2            string `cnab:"286,336"`
	Valor2           string `cnab:"336,536"`
	Brancos          string `cnab:"536,744"`
	NumeroSequencial int    `cnab:"744,750"`
}

type ShippingDetailCharge struct {
	TipoRegistro      int     `cnab:"0,1"`
	Identificador     string  `cnab:"1,36"`
	EmailDevedor      string  `cnab:"36,113"`
	LogradouroDevedor string  `cnab:"113,313"`
	CidadeDevedor     string  `cnab:"313,513"`
	EstadoDevedor     string  `cnab:"513,515"`
	CEPDevedor        string  `cnab:"515,523"`
	ModalidadeAbat    int     `cnab:"523,524"`
	ValorAbatimento   float64 `cnab:"524,541"`
	ModalidadeDesc    int     `cnab:"541,542"`
	DataDesconto1     int     `cnab:"542,550"`
	ValorDesconto1    float64 `cnab:"550,567"`
	DataDesconto2     int     `cnab:"567,575"`
	ValorDesconto2    float64 `cnab:"575,592"`
	DataDesconto3     int     `cnab:"592,600"`
	ValorDesconto3    float64 `cnab:"600,617"`
	ModalidadeJuros   int     `cnab:"617,618"`
	ValorJuros        float64 `cnab:"618,635"`
	ModalidadeMulta   int     `cnab:"635,636"`
	ValorMulta        float64 `cnab:"636,653"`
	Brancos           string  `cnab:"653,744"`
	NumeroSequencial  int     `cnab:"744,750"`
}

type ShippingTrailer struct {
	TipoRegistro     int     `cnab:"0,1"`
	Brancos          string  `cnab:"1,712"`
	ValorTotal       float64 `cnab:"712,729"`
	QtdeRegistros    int     `cnab:"729,744"`
	NumeroSequencial int     `cnab:"744,750"`
}

func (s *Shipping) FromFile(file io.Reader) error {
	var (
		shippingHeader2           ShippingHeader
		shippingDetail2           []ShippingDetail
		shippingDetailAdditional2 []ShippingDetailAdditional
		shippingDetailCharge2     []ShippingDetailCharge
		shippingTrailer2          ShippingTrailer
	)

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading file:", err)
		return err
	}

	if err = gocnab.Unmarshal(data, map[string]interface{}{
		"0": &shippingHeader2,
		"1": &shippingDetail2,
		"2": &shippingDetailAdditional2,
		"3": &shippingDetailCharge2,
		"9": &shippingTrailer2,
	}); err != nil {
		log.Println("Error unmarshalling file:", err)
		return err
	}

	s.Header = shippingHeader2
	s.Detail = shippingDetail2
	s.Additional = shippingDetailAdditional2
	s.Charge = shippingDetailCharge2
	s.Trailer = shippingTrailer2

	return nil
}
