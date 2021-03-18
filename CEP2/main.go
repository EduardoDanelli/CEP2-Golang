package main

import (
	"cepAPI/search"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/cep", search.CEP)

	if err := r.Run(":9000"); err != nil {
		log.Fatal(err.Error())
	}
}
