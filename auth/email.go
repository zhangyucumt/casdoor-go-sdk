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

type emailForm struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Sender    string   `json:"sender"`
	Receivers []string `json:"receivers"`
}

func SendEmail(title string, content string, sender string, receivers ...string) error {
	form := emailForm{
		Title:     title,
		Content:   content,
		Sender:    sender,
		Receivers: receivers,
	}
	postBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}

	resp, err := doPost("send-email", nil, postBytes, false)
	if err != nil {
		return err
	}

	if resp.Status != "ok" {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}

type EmailProviderForm struct {
	Owner       string
	Username    string
	Password    string
	Host        string
	Port        int
	Title       string
	Content     string
	ProviderUrl string
	Name        string
	DisplayName string
	Type        string
}

func NewEmailProvider(host string, port int, username string, password string, owner string) EmailProviderForm {
	str := genRandomString(6)
	name := "provider_" + str
	displayName := "New Provider - " + str
	form := EmailProviderForm{
		Owner:       owner,
		Name:        name,
		DisplayName: displayName,
		Host:        host,
		Port:        port,
		Username:    username,
		Password:    password,
		Type:        "Default",
		ProviderUrl: "https://github.com/organizations/xxx/settings/applications/1234567",
		Content:     "You have requested a verification code at Casdoor. Here is your code: %s, please enter in 5 minutes.",
		Title:       "Casdoor Verification Code",
	}
	return form
}

func CreateEmailProvider(form EmailProviderForm) (bool, error) {
	provider := Provider{
		Owner:        form.Owner,
		Name:         form.Name,
		DisplayName:  form.DisplayName,
		Host:         form.Host,
		Port:         form.Port,
		ClientId:     form.Username,
		ClientSecret: form.Password,
		Type:         form.Type,
		ProviderUrl:  form.ProviderUrl,
		Content:      form.Content,
		Title:        form.Title,
		CreatedTime:  time.Now().Format("2006-01-02T15:04:05+08:00"),
		Category:     "Email",
		Method:       "Normal",
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

func UpdateEmailProvider(provider Provider, form EmailProviderForm) (bool, error) {
	if form.Owner != "" {
		provider.Owner = form.Owner
	}
	if form.Name != "" {
		provider.Name = form.Name
	}
	if form.DisplayName != "" {
		provider.DisplayName = form.DisplayName
	}
	if form.Host != "" {
		provider.Host = form.Host
	}
	if form.Port != 0 {
		provider.Port = form.Port
	}
	if form.Username != "" {
		provider.ClientId = form.Username
	}
	if form.Password != "" {
		provider.ClientSecret = form.Password
	}
	if form.Type != "" {
		provider.Type = form.Type
	}
	if form.ProviderUrl != "" {
		provider.ProviderUrl = form.ProviderUrl
	}
	if form.Content != "" {
		provider.Content = form.Content
	}
	if form.Title != "" {
		provider.Title = form.Title
	}

	postBytes, err := json.Marshal(provider)
	if err != nil {
		return false, err
	}

	resp, err := doPost("update-provider", map[string]string{
		"id": provider.Owner + "/" + provider.Name,
	}, postBytes, false)
	if err != nil {
		return false, err
	}
	return resp.Data == "Affected", nil
}
