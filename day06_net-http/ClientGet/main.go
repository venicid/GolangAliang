package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	requestGet()
}

func requestGet() {
	apiUrl := "http://127.0.0.1:8005"

	data := url.Values{} // map[string][]string
	data.Set("name", "admin")

	fmt.Printf("data: %#v\n", data) // data: url.Values{"name":[]string{"admin"}}

	urlObj, _ := url.ParseRequestURI(apiUrl + "/req/get")
	urlObj.RawQuery = data.Encode()

	fmt.Printf("urlObj: %#v\n", urlObj)
	// u: &url.URL{Scheme:"http", Opaque:"", User:(*url.Userinfo)(nil),
	//Host:"127.0.0.1:8005", Path:"/req/get",
	//RawPath:"", ForceQuery:false, RawQuery:"name=admin", Fragment:"", RawFragment:""}

	response, _ := http.Get(urlObj.String())
	fmt.Printf("response: %#v\n", response)
	// response: &http.Response{Status:"200 OK", StatusCode:200, Proto:"HTTP/1.1",
	//ProtoMajor:1, ProtoMinor:1, Header:http.Header{"Content-Length":[]string{"18"},
	//"Content-Type":[]string{"text/plain; charset=utf-8"}, "Date":[]string{"Sat, 05 Nov 2022 05:45:14 GMT"}},
	//Body:(*http.bodyEOFSignal)(0xc00020c040), ContentLength:18, TransferEncoding:[]string(nil), Close:false, Uncompressed:false,
	//Trailer:http.Header(nil), Request:(*http.Request)(0xc000140000), TLS:(*tls.ConnectionState)(nil)}

	resultByte, _ := ioutil.ReadAll(response.Body) // Body io.ReadCloser
	fmt.Println(string(resultByte))

	type Response struct {
		Name2 string
	}

	var res Response
	json.Unmarshal(resultByte, &res)
	fmt.Printf("res: %#v\n", res)
	fmt.Println(res.Name2)

}

/*
请求路由为： http://127.0.0.1:8005/req/get?name=root
返回数据为： ["root"]
*/
