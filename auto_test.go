package miniprog

import (
	"fmt"
	"testing"
)

func TestCode2SessionRequest(t *testing.T) {
	req := &Code2SessionRequest{
		Appid:  "testid",
		Secret: "secret",
	}
	fmt.Println(URLEncode(req))
	s := NewServer()
	request := &GetAccessTokenRequest{}
	resp := &GetAccessTokenResponse{}
	if err := s.GetAccessToken(request, resp); err != nil {
		t.Error(err)
	}
}
