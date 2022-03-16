package library

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var snow_seed int64 = 0

func HeaderParse(h map[string]interface{}, deep int) (header map[string]string) {
	header = make(map[string]string, len(h))
	for hk, hv := range h {
		tmp := hv.(map[string]interface{})
		if _, ok := tmp[tmp["type"].(string)]; ok {
			switch tmp["type"] {
			case "default":
				header[hk] = tmp["default"].(string)
			// case "objectjson":
			// 	param[pk] = ParamParse(tmp)
			// case "arrayjson":
			// 	param[pk] = ParamParse(tmp)
			case "string":
				header[hk] = ParseString(tmp["string"].(string))
			case "int":
				header[hk] = strconv.Itoa(ParseInt(tmp["int"].(string)))
			case "datetime":
				header[hk] = ParseDatetime(tmp["datetime"].(string))
			case "timestamp":
				header[hk] = strconv.Itoa(int(time.Now().Unix()))
			default:
				panic("field：" + hk + ",错误的type：" + tmp["type"].(string))
			}
		}
	}
	return
}

func ParamParse(format map[string]interface{}, deep int) interface{} {
	if len(format) < 1 {
		return nil
	}
	var param interface{}
	switch format["type"].(string) {
	case "default":
		param = format["default"].(string)
	case "string":
		param = ParseString(format["string"].(string))
	case "int":
		param = ParseInt(format["int"].(string))
	case "array":
		param = ParamParse2(format["array"].(map[string]interface{}), deep)
	case "object":
		param = ParamParse2(format["object"].(map[string]interface{}), deep)
	default:
		panic("未知数据类型：" + format["type"].(string))
	}
	return param
}

func ParamParse2(p map[string]interface{}, deep int) (param map[string]interface{}) {
	if len(p) < 1 {
		return
	}
	param = make(map[string]interface{}, len(p))
	for pk, pv := range p {
		tmp := pv.(map[string]interface{})
		if _, ok := tmp["type"]; ok {
			if _, ok := tmp[tmp["type"].(string)]; ok {
				switch tmp["type"].(string) {
				case "default":
					param[pk] = tmp["default"]
				case "object":
					param[pk] = ParamParse2(tmp["object"].(map[string]interface{}), deep+1)
				case "array":
					param[pk] = ParseArray(tmp["array"].(map[string]interface{}), deep+1)
				// case "objectjson":
				// 	param[pk] = ParamParse(tmp)
				// case "arrayjson":
				// 	param[pk] = ParamParse(tmp)
				case "string":
					param[pk] = ParseString(tmp["string"].(string))
				case "int":
					param[pk] = ParseInt(tmp["int"].(string))
				case "datetime":
					param[pk] = ParseDatetime(tmp["datetime"].(string))
				case "timestamp":
					param[pk] = int(time.Now().Unix())
				default:
					panic("field：" + pk + ",错误的type：" + tmp["type"].(string))
				}
			}
		}
		// if _, ok := tmp[tmp["type"]]; ok {
		// 	param[pk] = tmp[tmp["type"]]
		// }
	}
	return
}

func ResultParse(format map[string]interface{}, data interface{}, deep int) interface{} {
	if len(format) < 1 {
		return nil
	}
	dataType := strings.Replace(fmt.Sprintf("%T", data), " ", "", -1)
	// if format["type"].(string) != dataType {
	// 	panic("返回数据格式和预定义返回数据格式不符；formatType：" + format["type"].(string) + "；dataType：" + dataType)
	// }
	var parseData interface{}
	var resData interface{}
	switch format["type"].(string) {
	case "default":
		if dataType != "string" {
			resData = "返回数据不符合预期；返回值：" + data.(string) + "；预期：" + format["default"].(string)
			// panic("返回数据格式和预定义返回数据格式不符；formatType：" + format["type"].(string) + "；dataType：" + dataType)
		} else {
			resData = "返回值：" + data.(string) + "；预期：" + format["default"].(string)
		}
	case "string":
		if dataType != "string" {
			resData = "返回数据不符合预期；返回值：" + data.(string) + "；预期：" + format["default"].(string)
			// panic("返回数据格式和预定义返回数据格式不符；formatType：" + format["type"].(string) + "；dataType：" + dataType)
		} else {
			resData = "返回值：" + data.(string) + "；预期：" + format["default"].(string)
		}
	case "array":
		if dataType != "[]interface{}" {
			errJson := json.Unmarshal([]byte(data.(string)), &parseData)
			if errJson != nil {
				panic(errJson.Error())
			}
			dataType = strings.Replace(fmt.Sprintf("%T", parseData), " ", "", -1)
			if dataType != "[]interface{}" {
				panic("返回数据格式和预定义返回数据格式不符；formatType：" + format["type"].(string) + "；dataType：" + dataType)
			}
		} else {
			parseData = data
		}
		resData = ResultParse2(format, parseData.(map[string]interface{}), deep)
	case "object":
		if dataType != "map[string]interface{}" {
			errJson := json.Unmarshal([]byte(data.(string)), &parseData)
			if errJson != nil {
				panic(errJson.Error())
			}
			dataType = strings.Replace(fmt.Sprintf("%T", parseData), " ", "", -1)
			if dataType != "map[string]interface{}" {
				panic("返回数据格式和预定义返回数据格式不符；formatType：" + format["type"].(string) + "；dataType：" + dataType)
			}
		} else {
			parseData = data
		}
		resData = ResultParse2(format, parseData.(map[string]interface{}), deep)
	default:
		panic("未知数据类型：" + format["type"].(string))
	}
	return resData
}

func ResultParse2(format, formatdata map[string]interface{}, deep int) (result map[string]interface{}) {
	result = make(map[string]interface{}, len(format[format["type"].(string)].(map[string]interface{})))
	for pk, pv := range format[format["type"].(string)].(map[string]interface{}) {
		tmp := pv.(map[string]interface{})
		if _, ok := tmp["type"]; ok {
			if _, ok := tmp[tmp["type"].(string)]; ok {
				switch tmp["type"].(string) {
				case "default":
					if tmp["default"] != formatdata[pk] {
						result[pk] = "返回数据不符合预期；返回值：" + formatdata[pk].(string) + "；预期：" + tmp["default"].(string)
					} else {
						result[pk] = "返回值：" + formatdata[pk].(string) + "；预期：" + tmp["default"].(string)
					}
				case "object":
					result[pk] = ResultParse2(tmp["object"].(map[string]interface{}), formatdata[pk].(map[string]interface{}), deep+1)
				case "array":
					result[pk] = ResultParse2(tmp["array"].(map[string]interface{}), formatdata[pk].(map[string]interface{}), deep+1)
				// case "objectjson":
				// 	param[pk] = ParamParse(tmp)
				// case "arrayjson":
				// 	param[pk] = ParamParse(tmp)
				case "string":
					result[pk] = ResultParseString(tmp["string"].(string), formatdata[pk].(string))
				case "int":
					result[pk] = ResultParseString(tmp["string"].(string), formatdata[pk].(string))
				// case "datetime":
				// 	result[pk] = ParseDatetime(tmp["datetime"].(string))
				// case "timestamp":
				// 	result[pk] = ResultParseString(tmp["string"].(string), formatdata[pk].(string))
				// 	result[pk] = int(time.Now().Unix())
				default:
					panic("field：" + pk + ",错误的type：" + tmp["type"].(string))
				}
			}
		}
		// if _, ok := tmp[tmp["type"]]; ok {
		// 	param[pk] = tmp[tmp["type"]]
		// }
	}
	return
}

func ParseDatetime(datetime string) string {
	if ok := strings.Contains(datetime, "Y"); ok {
		datetime = strings.Replace(datetime, "Y", "2006", 1)
	}
	if ok := strings.Contains(datetime, "m"); ok {
		datetime = strings.Replace(datetime, "m", "01", 1)
	}
	if ok := strings.Contains(datetime, "d"); ok {
		datetime = strings.Replace(datetime, "d", "02", 1)
	}
	if ok := strings.Contains(datetime, "H"); ok {
		datetime = strings.Replace(datetime, "H", "15", 1)
	}
	if ok := strings.Contains(datetime, "i"); ok {
		datetime = strings.Replace(datetime, "i", "04", 1)
	}
	if ok := strings.Contains(datetime, "s"); ok {
		datetime = strings.Replace(datetime, "s", "05", 1)
	}
	return time.Now().Format(datetime)
}

func ParseString(s string) string {
	format1 := strings.Split(s, "|")
	slen_max, slen_min := ParseFormat(format1)
	// var format2 = make(map[string]string, len(format1))
	return RandString(slen_max, slen_min)
}

func ParseInt(s string) int {
	format1 := strings.Split(s, "|")
	slen_max, slen_min := ParseFormat(format1)
	return RandInt(slen_max, slen_min)
}

func ParseFormat(format1 []string) (slen_max int, slen_min int) {
	// var format2 = make(map[string]string, len(format1))
	slen_max = 26
	slen_min = 0
	if len(format1) > 0 {
		for _, fv := range format1 {
			tmpfv := strings.Split(fv, "@")
			if tmpfv[0] == "len" {
				slen_max, _ = strconv.Atoi(tmpfv[1])
				slen_min = slen_max
				break
			} else if tmpfv[0] == "max" {
				slen_max, _ = strconv.Atoi(tmpfv[1])
			} else if tmpfv[0] == "min" {
				slen_min, _ = strconv.Atoi(tmpfv[1])
			}
		}
	}
	return
}

func ParseArray(obj map[string]interface{}, deep int) interface{} {
	len := RandInt(1, 1)
	arr2 := make([]interface{}, len)
	if len > 0 {
		for ak := range arr2 {
			arr := ParamParse(obj, deep+1)
			arr2[ak] = arr
		}
	}
	return arr2
}

func RandChi(len int) string {
	a := make([]rune, len)
	snow_seed++
	r := rand.New(rand.NewSource(time.Now().UnixNano() + snow_seed))
	for i := range a {
		// time.Sleep(time.Nanosecond)
		// rand.Seed(time.Now().UnixNano())
		a[i] = rune(19968 + r.Int63n(40869-19968))
	}
	return string(a)
}

func RandString(max, min int) string {
	snow_seed++
	r := rand.New(rand.NewSource(time.Now().UnixNano() + snow_seed))
	len := 0
	if max < 1 || max-min < 0 {
		len = 0
	} else if (max - min) == 0 {
		len = r.Intn(max + 1)
	} else {
		len = r.Intn(max-min+1) + min
	}
	if len < 1 {
		return ""
	}
	snow_seed++
	r2 := rand.New(rand.NewSource(time.Now().UnixNano() + snow_seed))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r2.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func RandInt(max, min int) int {
	snow_seed++
	r := rand.New(rand.NewSource(time.Now().UnixNano() + snow_seed))
	len := 0
	if max < 1 || max-min < 0 {
		len = 0
	} else if (max - min) == 0 {
		len = r.Intn(max + 1)
	} else {
		len = r.Intn(max-min+1) + min
	}

	if len < 1 {
		return 0
	}
	max_int := ""
	for i := 0; i < len; i++ {
		max_int += "9"
	}
	min_int := "1"
	if min > 1 {
		for i := 1; i < min; i++ {
			min_int += "0"
		}
	}

	snow_seed++
	r2 := rand.New(rand.NewSource(time.Now().UnixNano() + snow_seed))

	max2, _ := strconv.Atoi(max_int)
	min2, _ := strconv.Atoi(min_int)
	int2 := r2.Intn(max2-min2) + min2

	return int2
}

func ResultParseString(format, data string) string {
	r := ""
	format1 := strings.Split(format, "|")
	slen_max, slen_min := ParseFormat(format1)
	if len(data) < slen_min {
		r = r + "返回数据长度小于预期最小长度；"
	}
	if len(data) > slen_max {
		r = r + "返回数据长度大于预期最大长度；"
	}
	r = r + "返回值：" + data + "；预期：" + format
	// var format2 = make(map[string]string, len(format1))
	return r
}

func ResultParseArray(format map[string]interface{}, obj []interface{}, deep int) interface{} {
	obj_len := len(obj)
	arr2 := make([]interface{}, obj_len)
	if obj_len > 0 {
		for ok, ov := range obj {
			arr2[ok] = ResultParse(format, ov, deep)
			// arr2[ak] = arr
		}
	}
	return arr2
}
