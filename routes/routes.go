package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// crea una funcion SetupRouter

func SetupRouters(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	r.GET("/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hola %s", nombre)
	})

}
