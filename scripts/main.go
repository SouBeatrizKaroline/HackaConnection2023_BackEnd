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

func main() {
	h1 := ShippingHeader{
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

	c1 := []ShippingDetail{
		{
			TipoRegistro:          1,
			Identificador:         "ABC123XYZ456",
			TipoPessoaRecebedor:   1,
			CNPJRecebedor:         79772365000194,
			AgenciaRecebedor:      123,
			ContaRecebedor:        1234567890,
			TipoContaRecebedor:    "CC",
			ChavePix:              "79772365000194",
			TipoCobranca:          "4",
			CodOcorrencia:         5,
			TimestampExpiracao:    20230617164755,
			DataVencimento:        20230617,
			ValidadeAposVenc:      20230619,
			ValorOriginal:         1500.50,
			TipoPessoaDevedor:     1,
			CNPJDevedor:           96833332000126,
			NomeDevedor:           "MARIA SANTOS",
			SolicitacaoPagador:    "PAGAR ATÉ O VENCIMENTO",
			ExclusivoPSPRecebedor: "",
			Brancos:               "",
			NumeroSequencial:      2,
		},
	}

	d1 := []ShippingDetailAdditional{
		{
			TipoRegistro:     2,
			Identificador:    "ABCD1234567890",
			Nome:             "João da Silva",
			Valor:            "lembre-se de efetuar o pagamento até a data de vencimento para evitar atrasos.",
			Nome2:            "MARIA SOUZA",
			Valor2:           "Agradecemos pela preferência e estamos à disposição para auxiliá-lo(a) no que for necessário.",
			Brancos:          "",
			NumeroSequencial: 3,
		},
	}

	o1 := []ShippingDetailCharge{
		{
			TipoRegistro:      3,
			Identificador:     "NOTA 9",
			EmailDevedor:      "devedor@example.com",
			LogradouroDevedor: "Av. Principal, 789 - Bairro Central, Cidade Nova - Estado",
			CidadeDevedor:     "SAO PAULO",
			EstadoDevedor:     "SP",
			CEPDevedor:        "12345-678",
			ModalidadeAbat:    1,
			ValorAbatimento:   30,
			ModalidadeDesc:    1,
			DataDesconto1:     20230625,
			ValorDesconto1:    20,
			DataDesconto2:     20230625,
			ValorDesconto2:    10,
			DataDesconto3:     20230625,
			ValorDesconto3:    40,
			ModalidadeJuros:   1,
			ValorJuros:        2,
			ModalidadeMulta:   1,
			ValorMulta:        200,
			Brancos:           "",
			NumeroSequencial:  4,
		},
	}

	f1 := ShippingTrailer{
		TipoRegistro:     9,
		Brancos:          "",
		ValorTotal:       1500.75,
		QtdeRegistros:    5,
		NumeroSequencial: 5,
	}

	data, err := gocnab.Marshal750(h1, c1, d1, o1, f1)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	filePath := "./files/file.txt"
	// /Users/jc/Projects/github.com/julioc98/citi/main.go
	// "front:front@mudar123"

	// Call the function to write the data to the file
	err = writeFile(data, filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var h2 ShippingHeader
	var c2 []ShippingDetail
	var d2 []ShippingDetailAdditional
	var o2 []ShippingDetailCharge
	var f2 ShippingTrailer

	if err = gocnab.Unmarshal(data, map[string]interface{}{
		"0": &h2,
		"1": &c2,
		"2": &d2,
		"3": &o2,
		"9": &f2,
	}); err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("h1 == h2", h1 == h2)
	for i := range c1 {
		fmt.Println("c1[i] == c2[i]", c1[i] == c2[i])
	}
	for i := range d1 {
		fmt.Println("d1[i] == d2[i]", d1[i] == d2[i])
	}
	for i := range o1 {
		fmt.Println("o1[i] == o2[i]", o1[i] == o2[i])
	}
	fmt.Println("f1 == f2", f1 == f2)
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
