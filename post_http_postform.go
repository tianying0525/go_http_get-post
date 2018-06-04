package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	fmt.Println("Post Information: \n")
	form := url.Values{"name": {"Morpheus"}, "job": {"Leader"}}
	res, err := http.PostForm("https://reqres.in/api/users", form)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, res.Body)
	fmt.Println("\n")
}