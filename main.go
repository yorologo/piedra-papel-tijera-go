package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Crear una nueva instancia del router Gin
    router := gin.Default()

    // 1. Servir archivos estáticos (CSS, JS, imágenes)
    router.Static("/css", "./css")
    router.Static("/js", "./js")

    // 2. Ruta para la página principal
    router.GET("/", func(c *gin.Context) {
        c.File("index.html")
    })

    // 3. Ruta específica para el juego "/play"
    router.GET("/play", func(c *gin.Context) {
        // Aquí puedes manejar la lógica del juego si es necesario
        c.File("index.html") // Puedes cambiar esta lógica según sea necesario
    })

    // Iniciar el servidor en el puerto 8080
    router.Run(":8080")
}
