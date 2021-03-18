package search

import (
	"cepAPI/controllers"
	"strings"

	"github.com/gin-gonic/gin"
)

func CEP(c *gin.Context) {

	cep := c.Query("cep")
	cep = strings.ReplaceAll(cep, "-", "")
	cep = strings.ReplaceAll(cep, ".", "")

	// Verificando se o CEP é valido.
	if len(cep) != 8 {
		c.AbortWithStatusJSON(400, gin.H{"error": "CEP inválido"})
		return
	}

	// Tentando buscar o CEP no CepAberto
	data, erro := controllers.GetCepAberto(cep)
	if !erro {
		c.JSON(200, gin.H{"data": data})
		return
	}

	// Tentando buscar o CEP no ViaCep
	data2, erro2 := controllers.GetViaCep(cep)
	if !erro2 {
		c.JSON(200, gin.H{"data": data2})
		return
	}

	c.AbortWithStatusJSON(400, gin.H{"error": "CEP não encontrado"})

}
