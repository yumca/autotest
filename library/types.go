package library

type ResultData struct {
	Status    int         `json:"status"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp string      `json:"timestamp"`
}

//参数结构
type Param struct {
	Title  string      `json:"title"`
	Typ    string      `json:"typ"`
	Url    string      `json:"url"`
	Header interface{} `json:"header"`
	Param  interface{} `json:"param"`
	Result interface{} `json:"result"`
}
