package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/SmarticleLABS/maalomaal/easypaisa_lib/easypaisa_pkg"
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

	ep := &easypaisa_pkg.EASYPAISA_IMPL{}
	compListResp, err := ep.GetCompanyList()
	if nil != err {
		log.Println(err)
		panic(err)
	}

	fmt.Printf("Company List is:%+v\n", compListResp)
}
