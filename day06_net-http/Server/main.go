package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Name string `json:"name"`
}

//返回响应内容
type Response struct {
	Data    map[string]string `json:"data"`
	Message string            `json:"message"`
	Success bool              `json:"success"`
}

// 处理GET请求: http://127.0.0.1:8005/req/get?name=root
func dealGetRequestHandler(w http.ResponseWriter, r *http.Request) {
	// 获取请求的参数
	query := r.URL.Query()

	// 方式1：通过字典下标取值
	if len(query["name"]) > 0 {
		name := query["name"][0]
		fmt.Println("通过字典下标获取：", name)
	}

	// 方式2：使用Get方法，，如果没有值会返回空字符串
	name2 := query.Get("name")
	fmt.Println("通过get方式获取：", name2)

	data := make(map[string]string)
	data["name"] = name2 + " hello"

	res := Response{
		Data:    data,
		Message: "成功",
		Success: true,
	}

	//w.Write([]byte(string(json.Marshal(d)))) // 返回string
	json.NewEncoder(w).Encode(res) // 返回json数据
}

// 处理POST请求: http://127.0.0.1:8005/req/post {"name": "root"}
func dealPostRequestHandler(w http.ResponseWriter, r *http.Request) {
	// 请求体数据
	bodyContent, _ := ioutil.ReadAll(r.Body)
	strData := string(bodyContent)

	var request Request
	json.Unmarshal([]byte(strData), &request) // gin.ShouldBind
	fmt.Printf("body content:[%s]\n", string(bodyContent))

	data := make(map[string]string)
	data["name"] = request.Name + " hello"

	res := Response{
		Data:    data,
		Message: "成功",
		Success: true,
	}

	//w.Write(json.Marshal(res))

	json.NewEncoder(w).Encode(res)
}
func main() {
	http.HandleFunc("/req/post", dealPostRequestHandler)
	http.HandleFunc("/req/get", dealGetRequestHandler)

	http.ListenAndServe(":8005", nil)
	// 在golang中，你要构建一个web服务，必然要用到http.ListenAndServe
	// 第二个参数必须要有一个handler
}
