package library

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func HeaderParse(h map[string]interface{}) (header map[string]string) {
	header = make(map[string]string, len(h))
	for hk, hv := range h {
		tmp := hv.(map[string]string)
		if _, ok := tmp[tmp["type"]]; ok {
			header[hk] = tmp[tmp["type"]]
		}
	}
	return
}

func ParamParse(p map[string]interface{}) (param map[string]interface{}) {
	param = make(map[string]interface{}, len(p))
	for pk, pv := range p {
		tmp := pv.(map[string]interface{})
		switch tmp["type"].(string) {
		case "default":
			param[pk] = tmp["default"]
		case "object":
			param[pk] = ParamParse(tmp)
		case "array":
			param[pk] = ParamParse(tmp)
		// case "objectjson":
		// 	param[pk] = ParamParse(tmp)
		// case "arrayjson":
		// 	param[pk] = ParamParse(tmp)
		case "string":
			format1 := strings.Split(tmp["string"].(string), "|")
			var format2 map[string]string
			var slen_max = 1
			var slen_min = 1
			for _, fv := range format1 {
				tmpfv := strings.Split(fv, "@")
				if tmpfv[0] == "len" {
					slen_max, _ = strconv.Atoi(tmpfv[1])
					slen_min = slen_max
					break
				} else if tmpfv[0] == "max" {

				}
				format2[tmpfv[0]] = tmpfv[1]
			}
			RandString(slen)
			param[pk] = tmp["default"]
		case "int":
		case "datetime":
		case "timestamp":
		default:
			panic("field：" + pk + ",错误的type：" + tmp["type"].(string))
		}
		// if _, ok := tmp[tmp["type"]]; ok {
		// 	param[pk] = tmp[tmp["type"]]
		// }
	}
	return
}

func ResultParse(r interface{}) {
	switch r.(type) {
	case string:
		//add your operations
	case int8:
		//add your operations
	case int16:
		//add your operations
	case map[string]string:
	default:
		// return errors.New("no this type")
	}
}

func RandChi(len int) string {
	a := make([]rune, len)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range a {
		// time.Sleep(time.Nanosecond)
		// rand.Seed(time.Now().UnixNano())
		a[i] = rune(19968 + r.Int63n(40869-19968))
	}
	return string(a)
}

func RandString(max, min int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Intn(26)
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
