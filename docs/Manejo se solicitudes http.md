# Manejos de Solicitudes http

```go
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

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	r.GET("/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hola %s", nombre)
	})
	r.POST("/usuarios", func(c *gin.Context) {

		var nuevoUsuario Usuario
		if err := c.BindJSON(&nuevoUsuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error a decodificar el Json"})
			return
		}
		if nuevoUsuario.Nombre == "" || nuevoUsuario.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Faltan datos"})
			return
		}
		usuarios = append(usuarios, nuevoUsuario)

		c.JSON(http.StatusOK, gin.H{"message": "Usuario creado correctamente", "Datos del usuario": usuarios})

	})

}



```