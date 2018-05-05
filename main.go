package main

import (
	"fmt"
	"strconv"
)

const PORT = 8086

func main() {
	listenPort := ":" + strconv.Itoa(PORT)
	fmt.Printf("Fake listening on port: %s\n", listenPort)
	/*
		router := fasthttprouter.New()
		router.GET("/", index)
		router.GET("/hello/:name", hello)
		router.POST("/api/login", authenticateUser)
		//router.POST("/api/recipient", addRecipient)
		log.Fatal(fasthttp.ListenAndServe(listenPort, router.Handler))
	*/
}
