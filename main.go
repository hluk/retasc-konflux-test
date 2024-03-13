package main

import (
	"log"

	"github.com/hluk/retasc-konflux-test/api"
	_ "github.com/hluk/retasc-konflux-test/docs"
)

//	@title			ReTaSC Konflux Test
//	@version		1.0
//	@description	Proof of concept

//	@contact.name	API Support
//	@contact.url	https://github.com/hluk/retasc-konflux-test
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	router := api.SetupRouter()
	err := router.Run("0.0.0.0:8081")
	if err != nil {
		log.Fatal(err)
	}
}
