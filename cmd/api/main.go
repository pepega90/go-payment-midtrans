package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct{}

func main() {
	app := Config{}
	godotenv.Load(".env")
	app.SetupKeyAccess()
	r := gin.Default()
	r.POST("/charge", app.BeliLewatBANK)
	r.Run()
}
