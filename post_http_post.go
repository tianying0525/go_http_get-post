package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	Name string
	Job  string
}

type JSONUser struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

func NewJSONUser(user User) JSONUser {
	return JSONUser{
		user.Name,
		user.Job,
	}
}

func main() {
	fmt.Println("Post Information: \n")
	new_user := User{"Morphus", "Leader"}
	jsonValue, _ := json.Marshal(NewJSONUser(new_user))
	res, err := http.Post("https://reqres.in/api/users", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, res.Body)
	fmt.Println("\n")
}
