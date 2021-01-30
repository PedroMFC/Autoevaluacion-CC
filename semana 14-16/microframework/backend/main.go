package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	
	
	r.GET("/ping", func(c *gin.Context) {
		
		c.JSON(http.StatusOK, gin.H{ "message": "pong" })
	})

	r.GET("/holamundo", func(c *gin.Context) {
		jsonData := []byte(`{"msg":"this worked"}`)

    	c.Data(http.StatusOK, gin.MIMEJSON, jsonData)
	})

	r.Run("0.0.0.0:8080") // Se conecta a 0.0.0.0 para poder recibir peticiones de todas las interfaces 
}