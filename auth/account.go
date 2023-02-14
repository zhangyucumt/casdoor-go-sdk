package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
)

type HumanCheck struct {
	Type         string      `json:"type"`
	AppKey       string      `json:"appKey"`
	Scene        string      `json:"scene"`
	CaptchaId    string      `json:"captchaId"`
	CaptchaImage interface{} `json:"captchaImage"`
}

func GetHumanCheck() (*HumanCheck, error) {
	url := GetUrl("get-human-check", nil)
	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}
	var humanCheck *HumanCheck
	err = json.Unmarshal(bytes, &humanCheck)
	if err != nil {
		return nil, err
	}
	return humanCheck, nil
}

type SendVerificationCodeBody struct {
	Type           string `json:"type"`      // oneof Email Phone
	Dest           string `json:"dest"`      // 要发送至的号码
	CheckType      string `json:"checkType"` // 保留字段
	CheckId        string `json:"checkId"`   // 发送human check返回的id
	CheckKey       string `json:"checkKey"`  // 图片验证码
	CheckUser      string `json:"checkUser"` // 是否要检查用户是否登录 这里要设置成否
	OrganizationId string `json:"organizationId"`
}

func NewSendVerificationCodeBody(Type string, dest string, checkId string, checkKey string) *SendVerificationCodeBody {
	return &SendVerificationCodeBody{Type: Type, Dest: dest, CheckId: checkId, CheckKey: checkKey, CheckType: "captcha", CheckUser: "false", OrganizationId: "admin/" + authConfig.OrganizationName}
}

func SendVerificationCode(form *SendVerificationCodeBody) error {
	url := GetUrl("send-verification-code", nil)
	var err error
	data := url2.Values{"type": {form.Type}, "dest": {form.Dest}, "checkType": {form.CheckType}, "checkId": {form.CheckId}, "checkKey": {form.CheckKey}, "checkUser": {form.CheckUser}, "organizationId": {form.OrganizationId}}
	r, _ := http.PostForm(url, data)
	respByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var resp Response
	err = json.Unmarshal(respByte, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "ok" {
		return fmt.Errorf(resp.Msg)
	}
	return nil
}

type CheckVerificationCodeBody struct {
	Dest string `json:"dest"`
	Code string `json:"code"`
}

func NewCheckVerificationCodeBody(dest string, code string) *CheckVerificationCodeBody {
	return &CheckVerificationCodeBody{Dest: dest, Code: code}
}

func CheckVerificationCode(form *CheckVerificationCodeBody) error {
	postBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}
	resp, err := doPost("check-verification-code", nil, postBytes, false)
	if err != nil {
		return err
	}
	if resp.Status != "ok" {
		return fmt.Errorf(resp.Msg)
	}
	return nil
}
