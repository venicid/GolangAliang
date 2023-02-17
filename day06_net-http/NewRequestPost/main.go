package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	baseUrl := "http://127.0.0.1:8005/req/post"

	body := Request{}
	body.Name = "root"
	bodyByte, _ := json.Marshal(body)

	request, _ := http.NewRequest("POST", baseUrl, strings.NewReader(string(bodyByte)))
	//request, _ := http.NewRequest("POST", baseUrl, bytes.NewBuffer(bodyByte))

	headers := request.Header
	headers.Add("token", "d8cdcf8427e")
	headers.Set("Content-Type", "application/json")

	fmt.Printf("request: %#v\n", request)
	// request: &http.Request{Method:"GET", URL:(*url.URL)(0xc00013e000),
	//Proto:"HTTP/1.1", ProtoMajor:1, ProtoMinor:1,
	//Header:http.Header{"Content-Type":[]string{"application/json"},
	//"Token":[]string{"d8cdcf8427e"}}, Body:io.ReadCloser(nil),
	//GetBody:(func() (io.ReadCloser, error))(nil), ContentLength:0,
	//TransferEncoding:[]string(nil), Close:false, Host:"127.0.0.1:8005",
	//Form:url.Values(nil), PostForm:url.Values(nil), MultipartForm:(*multipart.Form)(nil),
	//Trailer:http.Header(nil), RemoteAddr:"", RequestURI:"", TLS:(*tls.ConnectionState)(nil),
	//Cancel:(<-chan struct {})(nil), Response:(*http.Response)(nil),
	//ctx:(*context.emptyCtx)(0xc00000a098)}

	client := &http.Client{}
	response, _ := client.Do(request)
	defer response.Body.Close()
	fmt.Printf("response: %#v\n", response)
	// response: &http.Response{Status:"200 OK", StatusCode:200, Proto:"HTTP/1.1", ProtoMajor:1,
	//ProtoMinor:1, Header:http.Header{"Content-Length":[]string{"65"},
	//"Content-Type":[]string{"text/plain; charset=utf-8"},
	//"Date":[]string{"Sat, 05 Nov 2022 06:44:21 GMT"}},
	//Body:(*http.bodyEOFSignal)(0xc0000a2040), ContentLength:65, TransferEncoding:[]string(nil),
	//Close:false, Uncompressed:false, Trailer:http.Header(nil), Request:(*http.Request)(0xc000140000),
	//TLS:(*tls.ConnectionState)(nil)}

	resultByte, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(resultByte))

	// 转为go中的opj
	responseOpj := Response{}
	json.Unmarshal(resultByte, &responseOpj)
	fmt.Printf("responseObj:%#v\n", responseOpj) // responseObj:main.Response{Data:map[string]string{"name":"test hello"}, Message:"成功", Success:true}
	fmt.Println(responseOpj.Message)

}

//返回响应内容
type Response struct {
	Data    map[string]string `json:"data"`
	Message string            `json:"message"`
	Success bool              `json:"success"`
}

type Request struct {
	Name string `json:"name"`
}
