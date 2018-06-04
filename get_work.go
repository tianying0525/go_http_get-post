package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//"reflect"
)

type Story struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

func main() {

	fmt.Println("Get Information: \n")
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		var story Story
		json.Unmarshal([]byte(contents), &story)
		fmt.Println(story)
		fmt.Println("\n")
	}

}
