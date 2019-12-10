package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/getTest",getRequest)
	//http.HandleFunc("/postTest",postRequest)
	http.HandleFunc("/postForm",postRequest2)
	http.HandleFunc("/getJson",getJsonData)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

// get请求
func getRequest(writer http.ResponseWriter, request *http.Request) {
	query:=request.URL.Query()
	// 方法1
	name:=query["name"][0]
	// 方法2
	name2:=query.Get("name")
	fmt.Print("获取到name2:"+name2)
	writer.Write([]byte(name+"一起学习Go Web编程吧"));
}
// post请求
// 请求头--application/json
func postRequest(w http.ResponseWriter,r *http.Request)  {

	decoder := json.NewDecoder(r.Body);

	var params map[string]string

	decoder.Decode(&params)

	// 打印获取的参数
	fmt.Print("username=%s,password=%s",params["username"],params["password"])

	w.Write([]byte("username:"+params["username"]+",password:"+params["password"]))
}
// post请求
// 请求头--application/x-www-form-urlencoded
func postRequest2(w http.ResponseWriter,r *http.Request)  {

	r.ParseForm()
	username:=r.Form.Get("username");
	password:=r.Form.Get("password")

	// 打印获取的参数
	fmt.Print("username=%s,password=%s",username,password)

	w.Write([]byte("username:"+username+",password:"+password))
}

// 返回的数据为json结构
type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data Person `json:"data"`
}

func getJsonData(writer http.ResponseWriter, request *http.Request)  {
	res := Response{
		0,
		"success",
		Person{
			"Jack",
			20,
		},
	}
	json.NewEncoder(writer).Encode(res)
}