package main

import (
	"net/http"
)

// struct Generated by: https://transform.tools/json-to-go
type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/cep", buscaCEP)
	http.ListenAndServe(":8080", nil)
}

func buscaCEP(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	if request.URL.Path != "/cep" {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{"message": "Not Found"}`))
		return
	}

	cep := request.URL.Query().Get("cep")
	if cep == "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "cep is required"}`))
		return
	}

	response.Write([]byte(`{"cep":"01001-000","logradouro":"Praça da Sé","complemento":"lado ímpar","bairro":"Sé","localidade":"São Paulo","uf":"SP","ibge":"3550308","gia":"1004","ddd":"11","siafi":"7107"}`))
}
