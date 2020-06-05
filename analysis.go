package miniprog

import "net/http"

// GetDailyRetainRequest 日留存请求参数
type GetDailyRetainRequest struct {
	AccessToken string `json:"access_token"`
	BeginDate   string `json:"begin_date"`
	EndDate     string `json:"end_date"`
}

// GetDailyRetainResponse 日留存返回值
type GetDailyRetainResponse struct {
	RefDate    string `json:"ref_date"`
	VisitUvNew []struct {
		Key   int `json:"key"`
		Value int `json:"value"`
	} `json:"visit_uv_new"`
	VisitUv []struct {
		Key   int `json:"key"`
		Value int `json:"value"`
	} `json:"visit_uv"`
}

// GetDailyRetain 获取用户访问小程序日留存
func (s *Server) GetDailyRetain(req *GetDailyRetainRequest, resp *GetDailyRetainResponse) error {
	request, err := http.NewRequest("POST", s.url+"/sns/jscode2session?access_token="+req.AccessToken, ToBody(req))
	if err != nil {
		return err
	}
	return s.Fetch(request, resp)
}
