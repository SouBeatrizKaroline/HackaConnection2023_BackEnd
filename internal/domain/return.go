package domain

import "github.com/julioc98/gocnab"

// Return is a struct that represents a CNAB 750 file (pt-BR Retorno)
type Return struct {
	Header           ReturnHeader
	Detail           []ReturnDetail
	DetailAdditional []ReturnDetailAdditional
	DetailGeneration []ReturnDetailTransactionGeneration
	Trailer          ReturnTrailer
}

type ReturnHeader struct {
	TipoRegistro             int    `cnab:"0,1"`
	CodigoRetorno            int    `cnab:"1,2"`
	LiteralRetorno           string `cnab:"2,9"`
	CodigoServico            int    `cnab:"9,11"`
	LiteralServico           string `cnab:"11,26"`
	ISPB                     string `cnab:"26,34"`
	TipoPessoaRecebedor      int    `cnab:"34,36"`
	CNPJRecebedor            int    `cnab:"36,50"`
	AgenciaRecebedor         int    `cnab:"50,54"`
	ContaRecebedor           int    `cnab:"54,74"`
	TipoContaRecebedor       string `cnab:"74,78"`
	ChavePix                 string `cnab:"78,155"`
	DataGeracao              int    `cnab:"155,163"`
	CodigoConvenio           string `cnab:"163,193"`
	ExclusivoPSP             string `cnab:"193,253"`
	NomeRecebedor            string `cnab:"253,353"`
	CodigosErro              string `cnab:"354,383"`
	Brancos                  string `cnab:"383,731"`
	NumeroSequencial         int    `cnab:"731,741"`
	VersaoArquivo            int    `cnab:"741,744"`
	NumeroSequencialRegistro int    `cnab:"744,750"`
}

type ReturnDetail struct {
	TipoRegistro           int     `cnab:"0,1"`
	Identificador          string  `cnab:"1,36"`
	TipoPessoaRecebedor    int     `cnab:"36,38"`
	CNPJRecebedor          int     `cnab:"38,52"`
	AgenciaRecebedor       int     `cnab:"52,56"`
	ContaRecebedor         int     `cnab:"56,76"`
	Tipo                   string  `cnab:"76,80"`
	ChavePix               string  `cnab:"80,157"`
	TipoCobranca           string  `cnab:"157,158"`
	CodigoMovimento        int     `cnab:"158,160"`
	TimestampExpiracao     int     `cnab:"160,174"`
	DataVencimento         int     `cnab:"174,182"`
	ValidadeAposVencimento int     `cnab:"182,186"`
	ValorOriginal          float64 `cnab:"186,203"`
	TipoPessoaDevedor      int     `cnab:"203,205"`
	CNPJDevedor            int     `cnab:"205,219"`
	NomeDevedor            string  `cnab:"219,359"`
	SolicitacaoPagador     string  `cnab:"359,499"`
	ExclusivoPSPRecebedor  string  `cnab:"499,559"`
	DataMovimento          int     `cnab:"559,567"`
	CodigosErro            string  `cnab:"567,597"`
	Revisao                int     `cnab:"597,601"`
	TarifaCobranca         float64 `cnab:"601,618"`
	Brancos                string  `cnab:"618,744"`
	NumeroSequencial       int     `cnab:"744,750"`
}

type ReturnDetailAdditional struct {
	TipoRegistro     int    `cnab:"1,1"`
	Identificador    string `cnab:"2,36"`
	Nome             string `cnab:"37,86"`
	Valor            string `cnab:"87,286"`
	NomeOpcional     string `cnab:"287,336"`
	ValorOpcional    string `cnab:"337,536"`
	Brancos          string `cnab:"537,744"`
	NumeroSequencial int    `cnab:"745,750"`
}

type ReturnDetailTransactionGeneration struct {
	TipoRegistro     int    `cnab:"1,1"`
	Identificador    string `cnab:"2,36"`
	ChavePix         string `cnab:"37,113"`
	CodigoMovimento  int    `cnab:"114,115"`
	DataMovimento    int    `cnab:"116,123"`
	EMVQrCode        string `cnab:"124,623"`
	LocationLink     string `cnab:"624,700"`
	Brancos          string `cnab:"701,744"`
	NumeroSequencial int    `cnab:"745,750"`
}

type ReturnDetailReceivement struct {
	TipoRegistro          int     `cnab:"1,1"`     // 9(01)
	Identificador         string  `cnab:"2,36"`    // X(35)
	ISPB                  string  `cnab:"37,44"`   // X(08)
	TipoPessoaRecebedor   int     `cnab:"45,46"`   // 9(02)
	CPFCNPJRecebedor      int     `cnab:"47,60"`   // 9(14)
	AgenciaRecebedor      int     `cnab:"61,64"`   // 9(04)
	ContaRecebedor        int     `cnab:"65,84"`   // 9(20)
	TipoContaRecebedor    string  `cnab:"85,88"`   // X(04)
	ChavePix              string  `cnab:"89,165"`  // X(77)
	TipoCobranca          string  `cnab:"166,166"` // X(01)
	CodigoMovimento       int     `cnab:"167,168"` // 9(02)
	DataMovimento         int     `cnab:"169,176"` // 9(08)
	DataVencimento        int     `cnab:"177,184"` // 9(08)
	TimestampPagamento    int     `cnab:"185,198"` // 9(14)
	ValorOriginal         float64 `cnab:"199,215"` // 9(15)V9(2)
	ValorJuros            float64 `cnab:"216,232"` // 9(15)V9(2)
	ValorMulta            float64 `cnab:"233,249"` // 9(15)V9(2)
	ValorAbatimento       float64 `cnab:"250,266"` // 9(15)V9(2)
	ValorDesconto         float64 `cnab:"267,283"` // 9(15)V9(2)
	ValorFinal            float64 `cnab:"284,300"` // 9(15)V9(2)
	ValorPago             float64 `cnab:"301,317"` // 9(15)V9(2)
	TipoPessoaDevedor     int     `cnab:"318,319"` // 9(02)
	CPFCNPJDevedor        int     `cnab:"320,333"` // 9(14)
	TipoPessoaPagador     int     `cnab:"334,335"` // 9(02)
	CPFCNPJPagador        int     `cnab:"336,349"` // 9(14)
	NomePagador           string  `cnab:"350,489"` // X(140)
	MensagemPagador       string  `cnab:"490,629"` // X(140)
	CodigoLiquidacao      string  `cnab:"630,631"` // X(02)
	EndToEndID            string  `cnab:"632,663"` // X(32)
	Revisao               int     `cnab:"664,667"` // 9(4)
	ExclusivoPSPRecebedor string  `cnab:"668,727"` // X(60)
	TarifaCobranca        float64 `cnab:"728,744"` // 9(15)V9(2)
	NumeroSequencial      int     `cnab:"745,750"` // 9(06)
}

type ReturnTrailer struct {
	TipoRegistro     int     `cnab:"1,1"`     // 9(01)
	CodigoRetorno    int     `cnab:"2,2"`     // 9(01)
	CodigoServico    int     `cnab:"3,4"`     // 9(02)
	ISPB             string  `cnab:"5,12"`    // X(08)
	CodigosErro      string  `cnab:"13,42"`   // X(30)
	Brancos          string  `cnab:"43,712"`  // X(670)
	ValorTotal       float64 `cnab:"713,729"` // 9(15)V9(2)
	QtdeDetalhes     int     `cnab:"730,744"` // 9(15)
	NumeroSequencial int     `cnab:"745,750"` // 9(06)
}

func (r *Return) ToFile() ([]byte, error) {
	return gocnab.Marshal750(r)
}
