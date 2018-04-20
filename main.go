package main

import (
	// "encoding/json"
	// "flag"
	"log"
	"net/http"
	"html/template"
	// "os"
	"reflect"
)

// Concurso struct
type Concurso struct {
	ProximoConcurso      string            `json:"proximoConcurso"`
	ConcursoAnterior     string            `json:"concursoAnterior"`
	Forward              string            `json:"forward"`
	Concurso 						 string 					 `json:"concuso"`
	Mensagens            []string          `json:"mensagens"`
	Datas                int64             `json:"data"`
	Resultado            string            `json:"resultado"`
	Ganhadores           int               `json:"ganhadores"`
	GanhadoresQuina      int               `json:"ganhadores_quina"`
	GanhadoresQuadra     int               `json:"ganhadores_quadra"`
	Valor                float64           `json:"valor"`
	ValorQuina           float64           `json:"valor_quina"`
	ValorQuadra          float64           `json:"valor_quadra"`
	Acumulado            int               `json:"acumulado"`
	ValorAcumulado       float64           `json:"valor_acumulado"`
	DataInclusao         int64             `json:"dtinclusao"`
	ProxFinalZero        string            `json:"prox_final_zero"`
	AcFinalZero          float64           `json:"ac_final_zero"`
	ProxConcursoFinal    string            `json:"proxConcursoFinal"`
	Observacao           string            `json:"observacao"`
	RowGUID              string            `json:"rowguid"`
	IcConferido          string            `json:"ic_conferido"`
	DeLocalSorteio       string            `json:"de_local_sorteio"`
	NoCidade             string            `json:"no_cidade"`
	SgUf                 string            `json:"sg_uf"`
	VrEstimativa         float64           `json:"vr_estimativa"`
	DtProximoConcurso    int64             `json:"dt_proximo_concurso"`
	VrAcumuladoEspecial  float64           `json:"vr_acumulado_especial"`
	VrArrecadado         float64           `json:"vr_arrecadado"`
	IcConcursoEspecial   bool              `json:"ic_concurso_especial"`
	Error                bool              `json:"error"`
	RateioProcessamento  bool              `json:"rateioProcessamento"`
	SorteioAcumulado     bool              `json:"sorteioAcumulado"`
	ConcursoEspecial     string            `json:"concursoEspecial"`
	GanhadoresPorUf      []GanhadoresPorUF `json:"ganhadoresPorUf"`
	ResultadoOrdenado    string            `json:"resultadoOrdenado"`
	DataStr              string            `json:"dataStr"`
	DtProximoConcursoStr string            `json:"dt_proximo_concursoStr"`
}

// GanhadoresPorUF struct
type GanhadoresPorUF struct {
	ProximoConcurso   string   `json:"proximoConcurso"`
	ConcursoAnterior  string   `json:"concursoAnterior"`
	Forward           string   `json:"forward"`
	Mensagens         []string `json:"mensagens"`
	CoLoteria         string   `json:"coLoteria"`
	NuConcurso        string   `json:"nuConcurso"`
	SgUf              string   `json:"sg_uf"`
	QtdGanhadores     int      `json:"qtGanhadores"`
	NoCidade          string   `json:"no_cidade"`
	Total             int      `json:"total"`
	IcCanalEletronico string   `json:"icCanalEletronico"`
}

// Node struct
type Node struct {
	ContactID  int
	EmployerID int
	FirstName  string
	MiddleName string
	LastName   string
}

func main() {
	// var concurso string
	// flag.StringVar(&concurso, "concurso", "", "Código do concurso/evento")
	// flag.Parse()

	// if concurso == "" {
	// 	log.Fatalln("Obrigatório passar o código do concurso/evento como parâmetro")
	// }

	// var url = "http://loterias.caixa.gov.br/wps/portal/loterias/landing/megasena/!ut/p/a1/04_Sj9CPykssy0xPLMnMz0vMAfGjzOLNDH0MPAzcDbwMPI0sDBxNXAOMwrzCjA0sjIEKIoEKnN0dPUzMfQwMDEwsjAw8XZw8XMwtfQ0MPM2I02-AAzgaENIfrh-FqsQ9wNnUwNHfxcnSwBgIDUyhCvA5EawAjxsKckMjDDI9FQE-F4ca/dl5/d5/L2dBISEvZ0FBIS9nQSEh/pw/Z7_HGK818G0KO6H80AU71KG7J0072/res/id=buscaResultado/c=cacheLevelPage/=/?&concurso=" + concurso

	// client := &http.Client{}
	// request, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cookie01 := http.Cookie{Name: "DigestTracker", Value: "AAABYuQiyEQ", Domain: "loterias.caixa.gov.br", Path: "/wps", HttpOnly: false, Secure: false}
	// cookie02 := http.Cookie{Name: "_ga", Value: "GA1.3.2107937287.1504898819", Domain: ".caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}
	// cookie03 := http.Cookie{Name: "ai_user", Value: "j0jz4|2018-04-20T13:34:43.274Z", Domain: "loterias.caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}
	// cookie04 := http.Cookie{Name: "_ga", Value: "GA1.4.2107937287.1504898819", Domain: ".loterias.caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}
	// cookie05 := http.Cookie{Name: "_gid", Value: "GA1.4.1452474901.1524231283", Domain: ".loterias.caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}
	// cookie06 := http.Cookie{Name: "_pk_ref.4.968f", Value: "%5B%22%22%2C%22%22%2C1524231284%2C%22https%3A%2F%2Fwww.google.com.br%2F%22%5D", Domain: ".caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}
	// cookie07 := http.Cookie{Name: "_pk_id.4.968f", Value: "cc6c921da841d7d6.1504898819.65.1524231293.1524231284.", Domain: ".caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}
	// cookie08 := http.Cookie{Name: "security", Value: "true", Domain: "loterias.caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}

	// request.AddCookie(&cookie01)
	// request.AddCookie(&cookie02)
	// request.AddCookie(&cookie03)
	// request.AddCookie(&cookie04)
	// request.AddCookie(&cookie05)
	// request.AddCookie(&cookie06)
	// request.AddCookie(&cookie07)
	// request.AddCookie(&cookie08)

	// response, err := client.Do(request)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// if response != nil {
	// 	concurso := new(Concurso)
	// 	json.NewDecoder(response.Body).Decode(concurso)
	// }

	// response.Body.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", renderTemplate)

	http.ListenAndServe(":3000", mux)
	// container := []Node{
	// 	{1, 12, "Accipiter", "ANisus", "Nisus"},
	// 	{2, 42, "Hello", "my", "World"},
	// }

	// // We create the template and register out template function
	// t := template.New("t").Funcs(templateFuncs)
	// t, err := t.Parse(htmlTemplate)
	// if err != nil {
	// 		panic(err)
	// }

	// err = t.Execute(os.Stdout, container)
	// if err != nil {
	// 		panic(err)
	// }
}

// In the template, we use rangeStruct to turn our struct values
// into a slice we can iterate over
var htmlTemplate = `
{{ .Usuario }}
<form method="post">
	<input type="text" name="concurso">
	<button>Pesquisar</button>
</form>
`

// var templateFuncs = template.FuncMap{"rangeStruct": RangeStructer}

// RangeStructer takes the first argument, which must be a struct, and
// returns the value of each field in a slice. It will return nil
// if there are no arguments or first argument is not a struct
func RangeStructer(args ...interface{}) []interface{} {
    if len(args) == 0 {
        return nil
    }

    v := reflect.ValueOf(args[0])
    if v.Kind() != reflect.Struct {
        return nil
    }

    out := make([]interface{}, v.NumField())
    for i := 0; i < v.NumField(); i++ {
        out[i] = v.Field(i).Interface()
    }

    return out
}

func renderTemplate(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	t, err := template.New("foo").Parse(htmlTemplate)

	if req.Method == "POST" {
		t.
	}

	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(res, nil)
}
