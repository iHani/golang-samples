// do this
//go get github.com/coocood/jas

/* example from github.com/coocood/jas
 *
 **/
package main

import "github.com/coocood/jas"
import "fmt"
import "net/http"

type Hello struct{}

func (*Hello) Get(ctx *jas.Context) { // `GET /v1/hello`
	ctx.Data = "hello world"
	//response: `{"data":"hello world","error":null}`
}
func main() {
	router := jas.NewRouter(new(Hello))
	router.BasePath = "/v1/"
	fmt.Println(router.HandledPaths(true))
	//output: `GET /v1/hello`
	http.Handle(router.BasePath, router)
	// I changed the port because I 8080 is in use
	http.ListenAndServe(":8081", nil)
}
