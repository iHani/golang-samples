// do this
//go get github.com/coocood/jas

/* example from github.com/coocood/jas
 * slightly modified
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

/*

github.com/coocood/jas
The MIT License

Copyright (c) 2013 Ewan Chou.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
