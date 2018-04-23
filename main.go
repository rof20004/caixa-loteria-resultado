package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gobuffalo/packr"
)

func main() {
	mux := http.NewServeMux()
	box := packr.NewBox("./pages")
	mux.Handle("/", http.FileServer(box))
	mux.HandleFunc("/search", search)
	http.ListenAndServe(":3000", mux)
}

func search(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		codigoConcurso := req.URL.Query().Get("concurso")

		var url = "http://loterias.caixa.gov.br/wps/portal/loterias/landing/megasena/!ut/p/a1/04_Sj9CPykssy0xPLMnMz0vMAfGjzOLNDH0MPAzcDbwMPI0sDBxNXAOMwrzCjA0sjIEKIoEKnN0dPUzMfQwMDEwsjAw8XZw8XMwtfQ0MPM2I02-AAzgaENIfrh-FqsQ9wNnUwNHfxcnSwBgIDUyhCvA5EawAjxsKckMjDDI9FQE-F4ca/dl5/d5/L2dBISEvZ0FBIS9nQSEh/pw/Z7_HGK818G0KO6H80AU71KG7J0072/res/id=buscaResultado/c=cacheLevelPage/=/?&concurso=" + codigoConcurso

		client := &http.Client{}
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalln(err)
		}

		cookie01 := http.Cookie{Name: "DigestTracker", Value: "AAABYuQiyEQ", Domain: "loterias.caixa.gov.br", Path: "/wps", HttpOnly: false, Secure: false}
		cookie02 := http.Cookie{Name: "_ga", Value: "GA1.3.2107937287.1504898819", Domain: ".caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}
		cookie03 := http.Cookie{Name: "ai_user", Value: "j0jz4|2018-04-20T13:34:43.274Z", Domain: "loterias.caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}
		cookie04 := http.Cookie{Name: "_ga", Value: "GA1.4.2107937287.1504898819", Domain: ".loterias.caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}
		cookie05 := http.Cookie{Name: "_gid", Value: "GA1.4.1452474901.1524231283", Domain: ".loterias.caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}
		cookie06 := http.Cookie{Name: "_pk_ref.4.968f", Value: "%5B%22%22%2C%22%22%2C1524231284%2C%22https%3A%2F%2Fwww.google.com.br%2F%22%5D", Domain: ".caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}
		cookie07 := http.Cookie{Name: "_pk_id.4.968f", Value: "cc6c921da841d7d6.1504898819.65.1524231293.1524231284.", Domain: ".caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}
		cookie08 := http.Cookie{Name: "security", Value: "true", Domain: "loterias.caixa.gov.br", Path: "/", HttpOnly: false, Secure: false}

		request.AddCookie(&cookie01)
		request.AddCookie(&cookie02)
		request.AddCookie(&cookie03)
		request.AddCookie(&cookie04)
		request.AddCookie(&cookie05)
		request.AddCookie(&cookie06)
		request.AddCookie(&cookie07)
		request.AddCookie(&cookie08)

		response, err := client.Do(request)
		if err != nil {
			log.Println(err)
		}
		defer response.Body.Close()

		concurso, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
		}

		res.Header().Set("Content-Type", "application/json;charset=utf-8")
		res.WriteHeader(http.StatusOK)
		res.Write(concurso)
	}
}
