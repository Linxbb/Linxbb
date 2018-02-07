package controller

import (
	"net/http"
	"log"
	"io/ioutil"
	"time"
	"strconv"
	"crypto/md5"
	"net/url"
	"encoding/hex"
	"encoding/json"
	"runtime/debug"
)


const(
	// 每月免费翻译200万字符！！！
	reqUrlB = "http://api.fanyi.baidu.com/api/trans/vip/translate"
	fromB = "auto"     //源语种
	toB = "zh"         //翻译后语种
	appidB = ""       // 你的appid
	passwordB = ""    //你是密钥
)
// 调用传入要翻译的字符串
func TranslationBaidu(req string)string {
	salt := strconv.Itoa(int(time.Now().Unix()))
	q := req
	h:= md5.New()
	h.Write([]byte(appidB+q+salt+passwordB))
	sign := hex.EncodeToString(h.Sum(nil))
	//request
	u, _:= url.Parse(reqUrlB)
	par := u.Query()
	par.Set("q",q)
	par.Set("from",fromB)
	par.Set("to",toB)
	par.Set("appid",appidB)
	par.Set("salt",salt)
	par.Set("sign",sign)
	u.RawQuery = par.Encode()
	res,err := http.Get(u.String())
	if err != nil{
		log.Println("request err:",err)
	}
	defer res.Body.Close()
	response,err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Println("read err:",err)
	}
	var result translate
	err = json.Unmarshal(response,&result)
	if err != nil{
		log.Println("marshal err:",err)
	}
	defer func() {
		if p := recover(); p != nil {
			log.Printf("panic recover! p: %v", p)
			debug.PrintStack()
		}
	}()
	return result.Trans[0]["dst"]
}

type translate struct {
	Trans []map[string]string  `json:"trans_result"`
}
