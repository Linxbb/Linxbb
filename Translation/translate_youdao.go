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

const (
	reqUrl   = "http://openapi.youdao.com/api"
	from     = "auto"
	to       = "zh-CHS"
	appKey   = ""   //你的appkey
	password = ""   //你的password
)
//调用 传入需要翻译的字符串
func TranslationYoudao(req string) string {
	salt := strconv.Itoa(int(time.Now().Unix()))
	q := req
	h := md5.New()
	h.Write([]byte(appKey + q + salt + password))
	sign := hex.EncodeToString(h.Sum(nil))
	//request
	u, _ := url.Parse(reqUrl)
	par := u.Query()
	par.Set("q", q)
	par.Set("from", from)
	par.Set("to", to)
	par.Set("appKey", appKey)
	par.Set("salt", salt)
	par.Set("sign", sign)
	u.RawQuery = par.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		log.Println("request err:", err)
	}
	defer res.Body.Close()
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("read err:", err)
	}
	var result translateYoudao
	err = json.Unmarshal(response, &result)
	if err != nil {
		log.Println("marshal err:", err)
	}
	defer func() {
		if p := recover(); p != nil {
			log.Printf("panic recover! p: %v", p)
			debug.PrintStack()
		}
	}()
	return result.Trans[0]
}

type translateYoudao struct {
	Trans []string `json:"translation"`
}
