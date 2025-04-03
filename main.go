package main

import (
	"gwt"
	"net/http"
)

func main() {
	gwt := gwt.New()
	gwt.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test...."))
	})
	gwt.Run(":8888")
}
