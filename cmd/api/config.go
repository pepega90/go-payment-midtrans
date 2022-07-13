package main

import (
	"os"

	"github.com/midtrans/midtrans-go"
)

func (app *Config) SetupKeyAccess() {
	midtrans.ServerKey = os.Getenv("SERVER_KEY")
	midtrans.Environment = midtrans.Sandbox
	midtrans.ClientKey = os.Getenv("CLIENT_KEY")
}
