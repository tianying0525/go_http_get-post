package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
	//"reflect"
)

type Data_Type struct {
	Id         int
	First_Name string
	Last_Name  string
	Avatar     string
}


type User struct {
	Data_Info Data_Type
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
		contents, err :=ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}


		var user User
		json.Unmarshal([]byte(contents), &user)
		var data_new Data_Type = user.Data_Info
		fmt.Println(data_new.First_Name)
		
	}

}
