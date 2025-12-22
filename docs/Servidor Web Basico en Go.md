# Servidor Basico en GO

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//c.String(200, "¡Hello, World!")
		c.JSON(200, gin.H{
			"message": "Hola Mundo",
		})
	})

	r.Run(":8080")

}




```