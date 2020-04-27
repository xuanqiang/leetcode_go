/**
 * $ curl http://localhost:9999/
 * $ curl http://localhost:9999/hello
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", indexHandler)
	//第二个参数代表处理所有的HTTP请求的实例
	//nil 代表使用标准库中的实例处理
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		//单引号围绕的字符字面值
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
