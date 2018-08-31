package main

import (
	"log"
	"net/http"
)

func request() {
	http.Get("http://localhost:1234/")
}

func Ddos() {
	n := 1000000
	i := 0
	for i < n {
		log.Printf("%d", i)
		go request()
		i++
	}
}
