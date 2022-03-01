package library

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpRequest(method string, url string, data map[string]string, header map[string]string) {
	// data = `{"type":"10","msg":"hello."}`
	if _, ok := data["get"]; ok {
		url = url + "/" + data["get"]
	}
	request, err := http.NewRequest(strings.ToUpper(method), url, strings.NewReader(data["post"]))
	if err != nil {
		panic("New post error:" + err.Error() + "\n")
	}
	//application/x-www-form-urlencoded
	contentType := "application/json"
	for hk, hv := range header {
		if hk == "content-type" {
			contentType = hv
		} else {
			request.Header.Set(hk, hv)
		}
	}
	//application/json
	request.Header.Set("Content-Type", contentType)
	//post数据并接收http响应
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("post data error:%v\n", err)
	} else {
		fmt.Println("post a data successful.")
		respBody, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("response data:%v\n", string(respBody))
	}
}
