package miniprog

import "net/http"

// BusinessLicenseRequest 营业执照识别参数
type BusinessLicenseRequest struct {
	AccessToken string `json:"access_token"`
	ImgURL      string `json:"img_url"`
}

// IdcardRequest 身份证识别参数
type IdcardRequest struct {
	AccessToken string `json:"access_token"`
	ImgURL      string `json:"img_url"`
}

// IdcardResponse 身份证识别返回值
type IdcardResponse struct {
	Errcode     string `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	ID          string `json:"id"`
	Addr        string `json:"addr"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

// BusinessLicenseResponse 营业执照识别返回值
type BusinessLicenseResponse struct {
	Errcode             int    `json:"errcode"`
	Errmsg              string `json:"errmsg"`
	RegNum              string `json:"reg_num"`
	Serial              string `json:"serial"`
	LegalRepresentative string `json:"legal_representative"`
	EnterpriseName      string `json:"enterprise_name"`
	TypeOfOrganization  string `json:"type_of_organization"`
	Address             string `json:"address"`
	TypeOfEnterprise    string `json:"type_of_enterprise"`
	BusinessScope       string `json:"business_scope"`
	RegisteredCapital   string `json:"registered_capital"`
	PaidInCapital       string `json:"paid_in_capital"`
	ValidPeriod         string `json:"valid_period"`
	RegisteredDate      string `json:"registered_date"`
}

// BusinessLicense 提供基于小程序的营业执照 OCR 识别
func (s *Server) BusinessLicense(req *BusinessLicenseRequest, resp *BusinessLicenseResponse) error {
	request, err := http.NewRequest("POST", s.url+"/cv/ocr/bizlicense?+"+URLEncode(req), nil)
	if err != nil {
		return err
	}
	return s.Fetch(request, resp)
}

// Idcard 提供基于小程序的身份证 OCR 识别
func (s *Server) Idcard(req *IdcardRequest, resp *IdcardResponse) error {
	request, err := http.NewRequest("POST", s.url+"/cv/ocr/idcard?+"+URLEncode(req), nil)
	if err != nil {
		return err
	}
	return s.Fetch(request, resp)
}
