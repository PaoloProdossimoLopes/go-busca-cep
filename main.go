package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// struct Generated by: https://transform.tools/json-to-go
type ViaCep struct {
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
	http.HandleFunc("/cep", buscaCEPHandler)
	http.ListenAndServe(":8080", nil)
}

func buscaCEPHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	if request.URL.Path != "/cep" {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{"message": "Not Found"}`))
		return
	}

	cep := request.URL.Query().Get("code")
	if cep == "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "cep is required"}`))
		return
	}

	viaCep, viaCepError := buscaCep(cep)
	if viaCepError != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "Internal Server Error"}`))
		return
	}

	json.NewEncoder(response).Encode(viaCep)
}

func buscaCep(cep string) (*ViaCep, error) {
	viaCepResponse, viaCepRequestError := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if viaCepRequestError != nil {
		return nil, viaCepRequestError
	}
	defer viaCepResponse.Body.Close()

	viaCepData, parseViaCepError := ioutil.ReadAll(viaCepResponse.Body)
	if parseViaCepError != nil {
		return nil, parseViaCepError
	}

	var viaCep ViaCep
	unmarshalViaCepError := json.Unmarshal(viaCepData, &viaCep)
	if unmarshalViaCepError != nil {
		return nil, unmarshalViaCepError
	}

	return &viaCep, nil
}
