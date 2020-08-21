package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reactgoboilerplate/router"
)

func main(){
	os.Setenv("TEST_ENV", "My test environment variable")

	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}