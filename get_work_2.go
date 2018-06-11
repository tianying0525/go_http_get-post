package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Info struct {
	Apple string
	Sky   string
}

func main() {
	url := "http://damp-headland-31920.herokuapp.com/test"
	content := contentFromServer(url)

	infos := infosFromJson(content)
	for _, info := range infos {
		fmt.Printf("APPLE is %v and SKY is %v\n", info.Apple, info.Sky)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {
	res, err := http.Get(url)
	checkError(err)

	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	checkError(err)

	return string(bytes)
}

func infosFromJson(content string) []Info {
	infos := make([]Info, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var info Info
	for decoder.More() {
		err := decoder.Decode(&info)
		checkError(err)
		infos = append(infos, info)
	}

	return infos
}
