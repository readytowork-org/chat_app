package main

import (
	"letschat/bootstrap"

	"go.uber.org/fx"
)

func main() {
	// infrastructure.InitializeLogger()
	// infrastructure.Logger.Info("started Server")
	fx.New(bootstrap.Module).Run()

}
