package main

import (
	"github.com/go_payment_midtrans/config"
	"github.com/go_payment_midtrans/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	config.SetupMidtransKeyAccess()
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/home", handlers.Index)
	r.POST("/charge", handlers.BeliLewatBANK)
	r.Run()
}
