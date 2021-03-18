package controllers

import (
	"cepAPI/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetViaCep(cep string) (models.Cidade, bool) {

	stop := false

	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		stop = true
		log.Fatal("Erro ao conectar com o ViaCEP")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		stop = true
		log.Fatal("Erro ao ler resposta")
	}

	defer resp.Body.Close()
	fmt.Printf("%s\n", data)

	var cidadeinfo models.CidadeViaCep
	if err := json.Unmarshal([]byte(data), &cidadeinfo); err != nil {
		stop = true
		log.Fatal("Erro ao conectar com ViaCEP")
	}

	if cidadeinfo.Erro {
		stop = true
	}

	return models.Cidade{
		Cep:        cidadeinfo.Cep,
		Logradouro: cidadeinfo.Logradouro,
		Bairro:     cidadeinfo.Bairro,
		Uf:         cidadeinfo.Uf,
		Ddd:        cidadeinfo.Ddd,
	}, stop
}
