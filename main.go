package main

import (
	"bitcoin-klever-api/routes"
	"os"
)

func main() {
	os.Setenv("USERNAME", "support")
	os.Setenv("PASSWORD", "Fg+GJKDACKIEOD3XVps=")
	os.Setenv("URL", "https://bitcoin.explorer.klever.io/api/v2/")

	routes.HandleRequest()
}
