package main

import (
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
    "tsacodingchallenge/config"
    "tsacodingchallenge/routes"
)

// @title API Title
// @version 1.0
// @description This is a sample API documentation
// @host localhost:8080
// @BasePath /
func main() {
    client, err := config.ConnectMongoDB()
    if err != nil {
        panic(err)
    }

    r := gin.Default()

    // Swagger documentation route
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Initialize your routes
    routes.InitializeRoutes(r, client)

    // Start the server
    r.Run(":8080")
}
