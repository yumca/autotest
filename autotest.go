package main

import (
	"autotest/library"
	"context"
	"encoding/json"
	"flag"
	"fmt"
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
		p_header := pv.Header.(map[string]interface{})
		header := library.HeaderParse(p_header, 1)
		fmt.Println("pk:", pk, "\nheader:", header)
		p_param := pv.Param.(map[string]interface{})
		post := library.ParamParse(p_param["post"].(map[string]interface{}), 1)
		if post != nil {
			jsonData, err := json.Marshal(post)
			if err != nil {
				panic(err.Error())
			}
			postString := string(jsonData)
			fmt.Println("pk:", pk, "\npost:", postString)
		}
		get := library.ParamParse(p_param["get"].(map[string]interface{}), 1)
		if get != nil {
			jsonData, err := json.Marshal(get)
			if err != nil {
				panic(err.Error())
			}
			getString := string(jsonData)
			fmt.Println("pk:", pk, "\nget:", getString)
		}
		fmt.Println("pk:", pk, "\nget:", get)
		// p_result := pv.Result.(map[string]interface{})
		// result := library.ParamParse(p_header, 1)
	}
}
