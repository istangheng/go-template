package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
	"time"
	"encoding/json"
	"bytes"
)

// 基本的GET请求
func fn1()  {
	resp, err := http.Get("http://httpbin.org/get")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    fmt.Println(string(body))

    if resp.StatusCode == 200 {
        fmt.Println("ok")
    }
}

// GET请求添加请求头
func fn2()  {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

    req,_ := http.NewRequest("GET","http://httpbin.org/get",nil)

    req.Header.Add("name","admin")
    req.Header.Add("age","18")

    resp,_ := client.Do(req)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Printf(string(body))

	if resp.StatusCode == 200 {
        fmt.Println("ok")
    }
}

// 基本的POST请求
func fn3()  {
	data := make(map[string]interface{})
    data["name"] = "admin"
    data["age"] = "18"

    bytesData, _ := json.Marshal(data)

    resp, _ := http.Post("http://httpbin.org/post","application/json", bytes.NewReader(bytesData))

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))

	if resp.StatusCode == 200 {
        fmt.Println("ok")
    }
}

// POST请求带参数
func fn4()  {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	data := `{"name":"admin", "age": "23"}`
    playload := bytes.NewBuffer([]byte(data))

    req, _ := http.NewRequest("POST","http://httpbin.org/post", playload)
    resp, _ := client.Do(req)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
  
func main() {
    fn4()
}