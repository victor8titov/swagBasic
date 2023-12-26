package main

import (
	"net/http"

	"github.com/swaggo/swag/example/basic/api"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	My description.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.email	tomato@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

func main() {
	http.HandleFunc("/testapi/get-string-by-int/", api.GetStringByInt)
	http.HandleFunc("/testapi/get-struct-array-by-string/", api.GetStructArrayByString)
	http.HandleFunc("/testapi/upload", api.Upload)
	http.HandleFunc("/testapi/get-info", api)
	http.ListenAndServe(":8080", nil)
}
