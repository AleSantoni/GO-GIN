package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

// crea una funcion SetupRouter

func SetupRouters(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Tilte":   "Mi Aplicacion",
			"Heading": "Hola Mundo",
			"Message": "Bienvenido a Mi Aplicacion Web con GIN",
		})
	})

	r.Static("/static", "./static")

}
