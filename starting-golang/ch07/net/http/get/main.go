package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
)

func main() {
	res, err1 := http.Get("https://google.com")
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println(res.StatusCode)
	fmt.Println(res.Proto)
	fmt.Println(res.Header)
	fmt.Println(res.Header["Content-Type"])
	fmt.Println(res.Request.Method)
	fmt.Println(res.Request.URL)

	defer res.Body.Close()

	body, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(string(body))
}
