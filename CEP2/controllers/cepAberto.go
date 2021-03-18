package controllers

import (
	"cepAPI/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetCepAberto(cep string) (models.Cidade, bool) {

	stop := false

	url := "https://www.cepaberto.com/api/v3/cep?cep=" + cep
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Token token=adee59a89420858065b72d17ce8c42b6")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		stop = true
		log.Printf("Erro: %v", err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		stop = true
		log.Fatal()
	}

	if string(data) == "{}" {
		stop = true
	}

	var cidadeInfoAberto models.CidadeAberto
	if err := json.Unmarshal([]byte(data), &cidadeInfoAberto); err != nil {
		stop = true
		log.Fatal("Erro ao conectar com ViaCEP")
	}

	return models.Cidade{
		Nome:       cidadeInfoAberto.Cidade.Nome,
		Cep:        cidadeInfoAberto.Cep,
		Logradouro: "",
		Bairro:     "",
		Uf:         cidadeInfoAberto.Estado.Sigla,
		Ddd:        cidadeInfoAberto.Ddd,
	}, stop

}
