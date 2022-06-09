package main

import (
	"go.uber.org/fx"
	"kloudlite.io/apps/dns/internal/framework"
)

func main() {
	module := framework.Module
	fx.New(module).Run()
}
