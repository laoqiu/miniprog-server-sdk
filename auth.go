package miniprog

import "net/http"

// Code2SessionRequest 登录凭证校验请求参数
type Code2SessionRequest struct {
	Appid     string `json:"appid"`
	Secret    string `json:"secret"`
	JsCode    string `json:"js_code"`
	GrantType string `json:"grant_type"`
}

// Code2SessionResponse 登录凭证校验返回值
type Code2SessionResponse struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

// GetPaidUnionIDRequest 获取用户支付后union请求参数
type GetPaidUnionIDRequest struct {
	AccessToken   string `json:"access_token"`
	Openid        string `json:"openid"`
	TransactionID string `json:"transaction_id"`
	MchID         string `json:"mch_id"`
	OutTradeNo    string `json:"out_trade_no"`
}

// GetPaidUnionIDResponse 获取用户支付后union返回值
type GetPaidUnionIDResponse struct {
	Unionid string `json:"unionid"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetAccessTokenRequest 获取accesstoken请求参数
type GetAccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Appid     string `json:"appid"`
	Secret    string `json:"secret"`
}

// GetAccessTokenResponse 获取accesstoken返回值
type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

// Code2Session 登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程
func (s *Server) Code2Session(req *Code2SessionRequest, resp *Code2SessionResponse) error {
	req.GrantType = "authorization_code"
	request, err := http.NewRequest("GET", s.url+"/sns/jscode2session?"+URLEncode(req), nil)
	if err != nil {
		return err
	}
	return s.Fetch(request, resp)
}

// GetPaidUnionID 用户支付完成后，获取该用户的 UnionID，无需用户授权
func (s *Server) GetPaidUnionID(req *GetPaidUnionIDRequest, resp *GetPaidUnionIDResponse) error {
	request, err := http.NewRequest("GET", s.url+"/wxa/getpaidunionid?"+URLEncode(req), nil)
	if err != nil {
		return err
	}
	return s.Fetch(request, resp)
}

// GetAccessToken 获取小程序全局唯一后台接口调用凭据（access_token）
func (s *Server) GetAccessToken(req *GetAccessTokenRequest, resp *GetAccessTokenResponse) error {
	req.GrantType = "client_credential"
	request, err := http.NewRequest("GET", s.url+"/cgi-bin/token?"+URLEncode(req), nil)
	if err != nil {
		return err
	}
	return s.Fetch(request, resp)
}
