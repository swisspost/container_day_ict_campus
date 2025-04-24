// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hello is a simple hello, world demonstration web server.
//
// It serves version information on /version and answers
// any other request like /name by saying "Hello, name!".
//
// See golang.org/x/example/outyet for a more sophisticated server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var greeting string
var addr string

func main() {
	// read greeting
	greeting = os.Getenv("GREETING")
	addr = os.Getenv("ADDR")

	// Server mux
	mux := http.NewServeMux()

	// Register handlers.
	mux.HandleFunc("/", greet)

	log.Printf("serving http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, logRequestMiddleware(mux)))
}

// logRequestMiddleware logs basic info of a HTTP request
// RemoteAddr: Network address that sent the request (IP:port)
// Proto: Protocol version
// Method: HTTP method
// URL: Request URL
// src: https://medium.com/geekculture/learn-go-middlewares-by-examples-da5dc4a3b9aa
func logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("LOG %s - %s %s %s\n", r.RemoteAddr, r.Proto, r.Method, r.URL)

		next.ServeHTTP(w, r)
	})
}

// src: https://github.com/golang/example/blob/master/helloserver/server.go
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<!DOCTYPE html>\n")
	fmt.Fprintf(w, "Hello, %s!\n", greeting)
}
