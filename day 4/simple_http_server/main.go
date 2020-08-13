package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func (writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello 🦄 @103cuong"))
	})
	http.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, `{"name":"Cuong Tran", "age": 22}`)
	})

	http.ListenAndServe(":5000", nil)
}
