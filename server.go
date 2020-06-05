package miniprog

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// DefaultEndpoint 微信服务端地址
var DefaultEndpoint = "https://api.weixin.qq.com"

// Server 定义服务端接口
type Server struct {
	url    string
	client http.Client
}

// NewServer 返回Server类
func NewServer() *Server {
	// client不验证https证书
	c := http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	return &Server{url: DefaultEndpoint, client: c}
}

// Fetch 请求微信服务器
func (s *Server) Fetch(req *http.Request, resp interface{}) error {
	res, err := s.client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, resp)
}
