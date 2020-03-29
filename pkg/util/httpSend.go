package util

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var (
	GET_METHOD    = "GET"
	POST_METHOD   = "POST"
	SENDTYPE_FORM = "form"
	SENDTYPE_JSON = "json"
)

type HttpSend struct {
	Url      string
	SendType string
	Header   map[string]string
	Body     map[string]string
	sync.RWMutex
}

func NewHttpSend(url string) *HttpSend {
	return &HttpSend{
		Url:      url,
		SendType: SENDTYPE_FORM,
	}
}

func (h *HttpSend) SetBody(body map[string]string) {
	h.Lock()
	defer h.Unlock()
	h.Body = body
}

func (h *HttpSend) SetHeader(Header map[string]string) {
	h.Lock()
	defer h.Unlock()
	h.Header = Header
}
func (h *HttpSend) SetSendType(sendType string) {
	h.Lock()
	defer h.Unlock()
	h.SendType = sendType
}

func (h *HttpSend) Get() ([]byte, error) {
	return h.send(GET_METHOD)
}

func (h *HttpSend) Post() ([]byte, error) {
	return h.send(POST_METHOD)
}

func GetUrlBuild(link string, data map[string]string) string {
	u, _ := url.Parse(link)
	q := u.Query()
	for k, v := range data {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (h *HttpSend) send(method string) ([]byte, error) {
	var (
		req       *http.Request
		resp      *http.Response
		client    http.Client
		send_data string
		err       error
	)
	if len(h.Body) > 0 {
		if strings.ToLower(h.SendType) == SENDTYPE_JSON {
			send_body, json_err := json.Marshal(h.Body)
			if json_err != nil {
				return nil, json_err
			}
			send_data = string(send_body)
		} else {
			send_body := http.Request{}
			send_body.ParseForm()
			for k, v := range h.Body {
				send_body.Form.Add(k, v)
			}
			send_data = send_body.Form.Encode()
		}
	}
	// 忽略https的证书
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, err = http.NewRequest(method, h.Url, strings.NewReader(send_data))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	//设置默认的header
	if len(h.Header) == 0 {
		//json
		if strings.ToLower(h.SendType) == SENDTYPE_JSON {
			h.Header = map[string]string{
				"Content-Type": "application/json; charset=utf-8",
			}
		} else {
			h.Header = map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			}
		}
	}

	for k, v := range h.Header {
		if strings.ToLower(k) == "host" {
			req.Host = v
		} else {
			req.Header.Add(k, v)
		}
	}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("error http code: %d", resp.StatusCode))
	}

	return ioutil.ReadAll(resp.Body)
}
