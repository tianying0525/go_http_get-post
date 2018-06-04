package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Story struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

/*
Use https://jsonplaceholder.typicode.com as test REST API to get different GET/POST
request
*/

func main() {
	//response, err := http.Get("https://damp-headland-31920.herokuapp.com/test")
	fmt.Println("Would you like to get or post? get/post")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if err != nil {
		fmt.Println(err)
	}

	switch text {

	case "get":
		fmt.Println("Get Is Pressed")
		fmt.Println("What do you want to get? Here are the attributes you could get (case sensitive):")
		fmt.Println("Id")
		fmt.Println("Title")
		fmt.Println("Body")

		reader := bufio.NewReader(os.Stdin)
		text_second, err := reader.ReadString('\n')
		text_second = strings.Replace(text_second, "\n", "", -1)

		if err != nil {
			fmt.Println(err)
		}

		response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
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
			var story []Story
			json.Unmarshal([]byte(contents), &story)

			switch text_second {
			case "Title":
				for i := 0; i < len(story); i++ {
					fmt.Printf("%v\n", story[i].Title)
				}
				break
			// implement whenver needed
			default:
				fmt.Printf("more coming: %v \n", text_second)
			}
		}
	case "post":
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			fmt.Println("Enter UserId")
			reader_userid := bufio.NewReader(os.Stdin)
			input_userid, err := reader_userid.ReadString('\n')
			input_userid = strings.Replace(input_userid, "\n", "", -1)

			fmt.Println("Enter Id")
			reader_id := bufio.NewReader(os.Stdin)
			input_id, err := reader_id.ReadString('\n')
			input_id = strings.Replace(input_id, "\n", "", -1)

			fmt.Println("Enter Title")
			reader_title := bufio.NewReader(os.Stdin)
			input_title, err := reader_title.ReadString('\n')
			input_title = strings.Replace(input_title, "\n", "", -1)

			fmt.Println("Enter Body")
			reader_body := bufio.NewReader(os.Stdin)
			input_body, err := reader_body.ReadString('\n')
			input_body = strings.Replace(input_body, "\n", "", -1)

			int_userid, _ := strconv.Atoi(input_userid)
			int_id, _ := strconv.Atoi(input_id)

			u := Story{UserId: int_userid, Id: int_id, Title: input_title, Body: input_body}
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(u)
			fmt.Printf("%v", u)
			res, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json; charset=utf-8", b)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			/*API would increment id number by 1, ignore. Test it by entering nothing at POST*/
			io.Copy(os.Stdout, res.Body)

			break
		}
	default:
		fmt.Println("Unkown Input")
		break
	}

}