package library

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/thinkeridea/go-extend/exnet"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var mylogs *MyLogs

/**
 * 去空
 */
func TrimEmpty(a []string) (ret []string) {
	aLen := len(a)
	for i := 0; i < aLen; i++ {
		if len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

/**
 * 日志
 */
func DoFileLog(filename, msg, types, mode string, data interface{}, httpreq *http.Request) (err error) {
	if httpreq == nil {
		httpreq = &http.Request{}
	}
	if mylogs == nil {
		mylogs, err = NewMyLogs(GetExecPath(), "", filename, ".log", httpreq)
		if err != nil {
			return
		}
	}
	mylogs.DoLogs(msg, types, mode, data)
	return
	//MyLogs.NewMyLogs()
	//$trace = debug_backtrace(DEBUG_BACKTRACE_IGNORE_ARGS, 1)
	//$this- > logs- > doLog($msg, $data, $type, current($trace))
}

/**
 * 日志
 */
func DoMyLogs(msg, types, mode string, data interface{}, httpreq *http.Request) (err error) {
	if httpreq == nil {
		httpreq = &http.Request{}
	}
	if mylogs == nil {
		mylogs, err = NewMyLogs(GetExecPath(), "", "", ".log", httpreq)
		if err != nil {
			return
		}
	}
	mylogs.DoLogs(msg, types, mode, data)
	return
	//MyLogs.NewMyLogs()
	//$trace = debug_backtrace(DEBUG_BACKTRACE_IGNORE_ARGS, 1)
	//$this- > logs- > doLog($msg, $data, $type, current($trace))
}

/**
 * 获取ip
 */
func Getip(r *http.Request) string {
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	if ip == "" {
		ip = "0.0.0.0"
	}
	return ip
}

func GetClientOs(r *http.Request) string {
	os := "other"
	userAgent := strings.ToLower(r.Header.Get("User-Agent"))
	if re := strings.IndexAny(userAgent, "iphone"); re != -1 {
		os = "iphone"
	} else if re := strings.IndexAny(userAgent, "android"); re != -1 {
		os = "android"
	} else if re := strings.IndexAny(userAgent, "micromessenger"); re != -1 {
		os = "weixin"
	} else if re := strings.IndexAny(userAgent, "ipad"); re != -1 {
		os = "ipad"
	} else if re := strings.IndexAny(userAgent, "ipod"); re != -1 {
		os = "ipod"
	} else if re := strings.IndexAny(userAgent, "windows nt"); re != -1 {
		os = "pc"
	}
	return os
}

//func GetItemId() string {
//    $hour = date("z") * 24 + date("H");
//    $hour = str_repeat("0", 4 - strlen($hour)) . $hour;
//    //	echo date("y") . $hour . PHP_EOL;
//    return date("y") . $hour . getRandNumber(10);
//}

func GetItemId() string {
	return ""
}

//返回毫秒时间戳 10+3
func GetMillisecond() string {
	// rand.Seed(time.Now().UnixNano())
	return strconv.FormatInt(time.Now().UnixMilli(), 10) //+ strconv.Itoa(rand.Intn(100))
}

//返回时间戳 10
func GetSecond() int {
	// rand.Seed(time.Now().UnixNano())
	return int(time.Now().Unix()) //+ strconv.Itoa(rand.Intn(100))
}

//设置api返回数据格式
func ApiResult(status int, msg string, data interface{}) ResultData {
	return ResultData{
		status,
		msg,
		data,
		strconv.FormatInt(time.Now().Unix(), 10),
	}
}

//生成32位md5字串
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// GBK 转 UTF-8
func GbkToUtf8(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), unicode.UTF8.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return s
	}
	return d
}

// UTF-8 转 GBK
func Utf8ToGbk(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		// log.Fatal(e)
		return s
	}
	return d
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home + "\\"
	}
	return os.Getenv("HOME") + "/"
}

func UserLocalLowDir() string {
	if runtime.GOOS == "windows" {
		path := UserHomeDir() + "AppData\\LocalLow\\"
		return path
	}
	return UserHomeDir() + "LocalLow/"
}

func UserLocalDir() string {
	if runtime.GOOS == "windows" {
		path := UserHomeDir() + "AppData\\Local\\"
		return path
	}
	return UserHomeDir() + "Local/"
}

func ProgramDir() string {
	conf, err := GetConf()
	if err != nil {
		return ""
	}
	path := ""
	if runtime.GOOS == "windows" {
		path = UserLocalLowDir() + conf.Setting.ServerName + "\\"
	} else {
		path = UserLocalLowDir() + conf.Setting.ServerName + "/"
	}
	if c := MkdirAll(path); !c {
		return ""
	}
	return path
}

func MkdirAll(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		//递归创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false
		}
	}
	return true
}

//获取当前文件执行路径
func GetExecPath() string {
	execFile, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(execFile)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}

func GetParamFileInfo(path string) (p []Param, err error) {
	if runtime.GOOS == "windows" {
		reg := regexp.MustCompile(`^[c-zC-Z](:\\)`)
		res := reg.FindAllString(path, -1)
		if len(res) < 1 {
			path = GetExecPath() + "\\" + path
		}
	} else {
		if !strings.HasPrefix(path, "/") {
			path = GetExecPath() + "/" + path
		}
	}
	file, osErr := os.Open(path)
	// 打开文件
	// file, osErr := os.Open("G:/WWW/golang/src/webchat/conf.json")
	if osErr != nil {
		err = errors.New("读取参数文件错误")
		return
	}
	// 关闭文件
	defer file.Close()
	var tmp []Param
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	errJson := decoder.Decode(&tmp)
	if errJson != nil {
		return tmp, errJson
	}
	p = tmp
	return
}
