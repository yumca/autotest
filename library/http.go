package library

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpRequest(method string, url string, data string, header map[string]string) ([]byte, error) {
	// data = `{"type":"10","msg":"hello."}`
	// if _, ok := data["get"]; ok {
	// 	url = url + "/" + data["get"]
	// }
	request, err := http.NewRequest(strings.ToUpper(method), url, strings.NewReader(data))
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
		return nil, err
	} else {
		if resp.StatusCode == 200 {
			respBody, _ := ioutil.ReadAll(resp.Body)
			return respBody, nil
		} else {
			return nil, errors.New("请求接口错误：" + resp.Status)
		}
	}
}
