package main

import (
	"autotest/library"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
)

var ctx context.Context

func main() {
	ctx = context.Background()
	// conf, err := library.GetConf()
	// if err != nil {
	// 	panic(err.Error())
	// }
	p := flag.String("p", "param.json", "参数结构文件")
	// c := flag.Int("c", 1, "并发数")
	// n := flag.Int("n", 10, "请求次数")
	// out := flag.String("o", "out.log", "返回数据输出文件")
	// library.DoFileLog("autotest", "mainServer GetConfig Error:"+err.Error(), "e", "main", "", nil)
	flag.Parse()
	param, err := library.GetParamFileInfo(*p)
	if err != nil {
		panic(err.Error())
	}
	// fmt.Printf("%+v\n", param)

	//根据参数结构生成数据
	for pk, pv := range param {
		url := pv.Url
		p_header := pv.Header.(map[string]interface{})
		header := library.HeaderParse(p_header, 1)
		fmt.Println("pk:", pk, "\nheader:", header)
		p_param := pv.Param.(map[string]interface{})
		postString := ""
		post := library.ParamParse(p_param["post"].(map[string]interface{}), 1)
		if post != nil {
			jsonData, err := json.Marshal(post)
			if err != nil {
				panic(err.Error())
			}
			postString = string(jsonData)
			fmt.Println("pk:", pk, "\npost:", postString)
		}
		get := library.ParamParse(p_param["get"].(map[string]interface{}), 1)
		getString := ""
		if get != nil {
			for gk, gv := range get {
				switch gv.(type) {
				case string:
					getString += gk + "=" + gv.(string)
				case int:
					getString += gk + "=" + strconv.Itoa(gv.(int))
				case map[string]interface{}:
					jsonData, err := json.Marshal(gv)
					if err != nil {
						panic(err.Error())
					}
					getString += gk + "=" + string(jsonData)
				}
			}
		}
		if getString != "" {
			url = url + "?" + getString
		}
		fmt.Println("pk:", pk, "\nurl:", url)
		//请求接口
		respBody, err := library.HttpRequest(pv.Typ, url, postString, header)
		if err != nil {
			panic(err.Error())
		}
		// fmt.Printf("response data:%v\n", string(respBody))
		var respData interface{}
		errJson := json.Unmarshal(respBody, &respData)
		if errJson != nil {
			panic(errJson.Error())
		}
		t := fmt.Sprintf("%#v", respData)
		fmt.Println(t)
		// library.ResultParse(pv.Result, respData)
		// result :=
		// p_result := pv.Result.(map[string]interface{})
		// result := library.ParamParse(p_header, 1)
	}
}
