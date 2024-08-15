package main

import (
	"consume-api-try/internal/provider/routes"
)

func main() {
	routes.SetupRoutes().Run(":8081")
}

