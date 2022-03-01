package main

import (
	"autotest/library"
	"context"
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
	fmt.Printf("%+v", param)
	//根据参数结构生成数据
	// for pk, pv := range param {

	// }
}
