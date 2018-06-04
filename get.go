package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//"reflect"
)

type User struct {
	Data_Info Data_Type `json:"data"`
}

type Data_Type struct {
	Id         int    `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Avatar     string `json:"avatar"`
}

func main() {

	fmt.Println("Get Information: \n")
	res, err := http.Get("https://reqres.in/api/users/2")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer res.Body.Close()
		//contents, err := ioutil.ReadAll(res.Body)
		contents, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		var user User
		json.Unmarshal([]byte(contents), &user)
		var data_new Data_Type = user.Data_Info
		fmt.Println(data_new)

		fmt.Println("\n")

	}

}
