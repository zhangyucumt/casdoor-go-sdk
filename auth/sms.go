// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"encoding/json"
	"fmt"
	"time"
)

type smsForm struct {
	Content   string   `json:"content"`
	Receivers []string `json:"receivers"`
}

func SendSms(content string, receivers ...string) error {
	form := smsForm{
		Content:   content,
		Receivers: receivers,
	}
	postBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}

	resp, err := doPost("send-sms", nil, postBytes, false)
	if err != nil {
		return err
	}

	if resp.Status != "ok" {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}

type SmsProviderForm struct {
	Owner        string
	ProviderUrl  string
	Name         string
	DisplayName  string
	Type         string
	ClientId     string
	ClientSecret string
	SignName     string
	TemplateCode string
	AppId        string
	Metadata     string
}

func NewSmsProvider(_type, clientId, clientSecret, signName, templateCode, appId, owner string) SmsProviderForm {
	str := genRandomString(6)
	name := "provider_" + str
	displayName := "New Provider - " + str
	form := SmsProviderForm{
		Owner:        owner,
		Name:         name,
		DisplayName:  displayName,
		Type:         _type,
		ProviderUrl:  "https://github.com/organizations/xxx/settings/applications/1234567",
		ClientId:     clientId,
		ClientSecret: clientSecret,
		SignName:     signName,
		TemplateCode: templateCode,
		AppId:        appId,
	}
	return form
}

func CreateSmsProvider(form SmsProviderForm) (bool, error) {
	provider := Provider{
		Owner:        form.Owner,
		Name:         form.Name,
		DisplayName:  form.DisplayName,
		ClientId:     form.ClientId,
		ClientSecret: form.ClientSecret,
		Type:         form.Type,
		ProviderUrl:  form.ProviderUrl,
		CreatedTime:  time.Now().Format("2006-01-02T15:04:05+08:00"),
		Category:     "SMS",
		Method:       "Normal",
		SignName:     form.SignName,
		TemplateCode: form.TemplateCode,
		AppId:        form.AppId,
		Metadata:     form.Metadata,
	}

	postBytes, err := json.Marshal(provider)
	if err != nil {
		return false, err
	}

	resp, err := doPost("add-provider", nil, postBytes, false)
	if err != nil {
		return false, err
	}
	return resp.Data == "Affected", nil
}
