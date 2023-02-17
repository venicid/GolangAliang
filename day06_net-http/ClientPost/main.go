package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	requestPost()
}
func requestPost() {
	url := "http://127.0.0.1:8005/req/post"

	// 表单数据 contentType := "application/x-www-form-urlencoded"
	contentType := "application/json"
	data := `{"name":"rootPort"}`

	resp, _ := http.Post(url, contentType, strings.NewReader(data))

	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}
