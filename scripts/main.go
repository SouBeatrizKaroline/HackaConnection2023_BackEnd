package main

import (
	"fmt"
	"os"

	"github.com/julioc98/gocnab"
)

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

// start -------------  return  -------------------- //
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

func main() {
	shippingHeader := ShippingHeader{
		TipoRegistro:             0,
		Operacao:                 1,
		LiteralRemessa:           "REMESSA",
		CodigoServico:            2,
		LiteralServico:           "PIX",
		ISPB:                     "12345678",
		TipoPessoa:               2,
		CNPJ:                     11486115000154,
		Agencia:                  1234,
		Conta:                    12345678,
		TipoConta:                "cc",
		ChavePix:                 "12345678000123",
		DataGeracao:              20230617,
		CodigoConvenio:           "ABC123456789",
		ExclusivoPSP:             "XYZ987654321",
		NomeRecebedor:            "Empresa ABC Ltda",
		Brancos:                  "",
		NumeroSequencial:         123,
		VersaoLayout:             2,
		NumeroSequencialRegistro: 1,
	}

	shippingDetail := []ShippingDetail{
		{
			TipoRegistro:          1,
			Identificador:         "ABC123XYZ456",
			TipoPessoaRecebedor:   2,
			CNPJRecebedor:         79772365000194,
			AgenciaRecebedor:      123,
			ContaRecebedor:        1234567890,
			TipoContaRecebedor:    "CC",
			ChavePix:              "79772365000194",
			TipoCobranca:          "4",
			CodOcorrencia:         5,
			TimestampExpiracao:    20230617164755,
			DataVencimento:        0,
			ValidadeAposVenc:      0,
			ValorOriginal:         0,
			TipoPessoaDevedor:     2,
			CNPJDevedor:           96833332000126,
			NomeDevedor:           "MARIA SANTOS",
			SolicitacaoPagador:    "PAGAR",
			ExclusivoPSPRecebedor: "",
			Brancos:               "",
			NumeroSequencial:      1,
		},
	}

	shippingDetailAdditional := []ShippingDetailAdditional{
		{
			TipoRegistro:     2,
			Identificador:    "ABC123XYZ456",
			Nome:             "João da Silva",
			Valor:            "lembre-se de efetuar o pagamento até a data de vencimento para evitar atrasos.",
			Nome2:            "MARIA SOUZA",
			Valor2:           "Agradecemos pela preferência e estamos à disposição para auxiliá-lo(a) no que for necessário.",
			Brancos:          "",
			NumeroSequencial: 1,
		},
	}

	shippingDetailCharge := []ShippingDetailCharge{
		{
			TipoRegistro:      3,
			Identificador:     "ABC123XYZ456",
			EmailDevedor:      "devedor@example.com",
			LogradouroDevedor: "Av. Principal, 789 - Bairro Central, Cidade Nova - Estado",
			CidadeDevedor:     "SAO PAULO",
			EstadoDevedor:     "SP",
			CEPDevedor:        "12345-678",
			ModalidadeAbat:    1,
			ValorAbatimento:   0,
			ModalidadeDesc:    1,
			DataDesconto1:     0,
			ValorDesconto1:    0,
			DataDesconto2:     0,
			ValorDesconto2:    0,
			DataDesconto3:     0,
			ValorDesconto3:    0,
			ModalidadeJuros:   1,
			ValorJuros:        2,
			ModalidadeMulta:   1,
			ValorMulta:        0,
			Brancos:           "",
			NumeroSequencial:  1,
		},
	}

	shippingTrailer := ShippingTrailer{
		TipoRegistro:     9,
		Brancos:          "",
		ValorTotal:       0,
		QtdeRegistros:    5,
		NumeroSequencial: 1,
	}

	data, err := gocnab.Marshal750(shippingHeader, shippingDetail, shippingDetailAdditional, shippingDetailCharge, shippingTrailer)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	filePath := "./file.txt"

	// Call the function to write the data to the file
	err = writeFile(data, filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var shippingHeader2 ShippingHeader
	var shippingDetail2 []ShippingDetail
	var shippingDetailAdditional2 []ShippingDetailAdditional
	var shippingDetailCharge2 []ShippingDetailCharge
	var shippingTrailer2 ShippingTrailer

	if err = gocnab.Unmarshal(data, map[string]interface{}{
		"0": &shippingHeader2,
		"1": &shippingDetail2,
		"2": &shippingDetailAdditional2,
		"3": &shippingDetailCharge2,
		"9": &shippingTrailer2,
	}); err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("h1 == h2", shippingHeader == shippingHeader2)
	for i := range shippingDetail {
		fmt.Println("shippingDetail[i] == shippingDetail2[i]", shippingDetail[i] == shippingDetail2[i])
	}
	for i := range shippingDetailAdditional {
		fmt.Println("shippingDetailAdditional[i] == shippingDetailAdditional2[i]", shippingDetailAdditional[i] == shippingDetailAdditional2[i])
	}
	for i := range shippingDetailCharge {
		fmt.Println("shippingDetailCharge[i] == shippingDetailCharge2[i]", shippingDetailCharge[i] == shippingDetailCharge2[i])
	}
	fmt.Println("shippingTrailer == shippingTrailer2", shippingTrailer == shippingTrailer2)

}

func writeFile(data []byte, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
