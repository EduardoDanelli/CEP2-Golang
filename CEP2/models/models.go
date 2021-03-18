package models

type CidadeViaCep struct {
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
	Erro        bool   `json:"erro"`
}

type CidadeAberto struct {
	Cidade    CidadeCidadeAberto `json:"cidade"`
	Estado    Estado             `json:"estado"`
	Altitude  float64            `json:"altitude"`
	Cep       string             `json:"cep"`
	Latitude  string             `json:"latitude"`
	Longitude string             `json:"longitude"`
	Ddd       string             `json:"ddd"`
}
type CidadeCidadeAberto struct {
	Ddd  int    `json:"ddd"`
	Nome string `json:"nome"`
}
type Estado struct {
	Sigla string `json:"sigla"`
}

type Cidade struct {
	Nome       string `json:"nome"`
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Uf         string `json:"uf"`
	Ddd        string `json:"ddd"`
}
