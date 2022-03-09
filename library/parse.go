package library

import (
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

func ParamParse(p map[string]interface{}, deep int) (param map[string]interface{}) {
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
					param[pk] = ParamParse(tmp["object"].(map[string]interface{}), deep+1)
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

func ResultParse(format, data interface{}) {
	switch format.(type) {
	case string:
		//add your operations
	case int8:
		//add your operations
	case int16:
		//add your operations
	case map[string]string:
	case map[string]interface{}:
	default:
		// return errors.New("no this type")
	}
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
			arr := ParamParse(obj[obj["type"].(string)].(map[string]interface{}), deep)
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
