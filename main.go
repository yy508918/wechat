package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	w.Write([]byte("hello"))
   	 io.WriteString(w, "hello io.writestring")
   	 fmt.Fprintf(w, "abc")

	log.Fatal(http.ListenAndServe(":80", nil))
}
