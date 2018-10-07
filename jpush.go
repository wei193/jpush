package jpush

import (
	"log"
	"net/http/httputil"

	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//地址变量
const (
	PushURL = "/v3/push"
)

//JPush JPush
type JPush struct {
	AppKey        string
	MasterSecret  string
	Authorization string
	ServerAddr    string
	DevMode       bool
}

//NewJPush NewJPush
func NewJPush(ServerAddr, AppKey, MasterSecret, mode string) (jpush *JPush) {
	if ServerAddr == "" {
		ServerAddr = "https://api.jpush.cn"
	}
	devmode := false
	if mode == "dev" {
		devmode = true
	}
	return &JPush{
		ServerAddr:    ServerAddr,
		AppKey:        AppKey,
		MasterSecret:  MasterSecret,
		Authorization: "Basic " + base64.StdEncoding.EncodeToString([]byte(AppKey+":"+MasterSecret)),
		DevMode:       devmode,
	}
}

func (j *JPush) requset(req *http.Request) ([]byte, error) {
	client := &http.Client{Timeout: 60 * time.Second}
	req.Header.Set("Authorization", j.Authorization)
	b, _ := httputil.DumpRequest(req, true)
	log.Println(string(b))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//createRequset 生成请求requset参数
func createRequset(surl, method string, p map[string]string, d interface{}) (req *http.Request, err error) {
	buf, err := json.Marshal(d)
	if d != nil && err != nil {
		return nil, err
	}

	val := make(url.Values)
	for k, v := range p {
		val.Add(k, v)
	}
	ContentType := ""
	switch method {
	case "GET":
		if strings.Index(surl, "?") != -1 {
			surl += "&" + val.Encode()
		} else {
			surl += "?" + val.Encode()
		}
	case "POST":
		if d == nil {
			buf = []byte(val.Encode())
			ContentType = "application/x-www-form-urlencoded"
		} else {
			ContentType = "application/json"
			if strings.Index(surl, "?") != -1 {
				surl += "&" + val.Encode()
			} else {
				surl += "?" + val.Encode()
			}
		}
	default:
		if strings.Index(surl, "?") != -1 {
			surl += "&" + val.Encode()
		} else {
			surl += "?" + val.Encode()
		}
	}

	req, err = http.NewRequest(method, surl, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	if ContentType != "" {
		req.Header.Add("Content-Type", ContentType)
	}
	return req, err
}
