package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ejemplo de respuesta en localhost:8080/ping
// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

//ejemplo con el nombre ingresado en la URL
func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) { //:name es el pathparam y gin lo utiliza para crear el contexto
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s como va todo?", name) //utilizo la libreria http.
	})

	router.Run(":8080")
}
