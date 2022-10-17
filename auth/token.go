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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/oauth2"
)

type Token struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	Application  string `xorm:"varchar(100)" json:"application"`
	Organization string `xorm:"varchar(100)" json:"organization"`
	User         string `xorm:"varchar(100)" json:"user"`

	Code          string `xorm:"varchar(100)" json:"code"`
	AccessToken   string `xorm:"mediumtext" json:"accessToken"`
	RefreshToken  string `xorm:"mediumtext" json:"refreshToken"`
	ExpiresIn     int    `json:"expiresIn"`
	Scope         string `xorm:"varchar(100)" json:"scope"`
	TokenType     string `xorm:"varchar(100)" json:"tokenType"`
	CodeChallenge string `xorm:"varchar(100)" json:"codeChallenge"`
	CodeIsUsed    bool   `json:"codeIsUsed"`
	CodeExpireIn  int64  `json:"codeExpireIn"`
}

func GetTokens() ([]*Token, error) {
	queryMap := map[string]string{
		"owner":     authConfig.OrganizationName,
		"tokenType": "AccessToken",
	}

	url := GetUrl("get-tokens", queryMap)

	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var tokens []*Token
	err = json.Unmarshal(bytes, &tokens)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func GetToken(name string) (*Token, error) {
	queryMap := map[string]string{
		"id":        fmt.Sprintf("%s/%s", authConfig.OrganizationName, name),
		"tokenType": "AccessToken",
	}

	url := GetUrl("get-token", queryMap)

	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var token *Token
	err = json.Unmarshal(bytes, &token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func UpdateToken(token *Token) (bool, error) {
	token.TokenType = "AccessToken"
	_, affected, err := modifyToken("update-token", token, nil)
	return affected, err
}

func AddToken(token *Token) (bool, error) {
	_, affected, err := modifyToken("add-token", token, nil)
	return affected, err
}

func DeleteToken(token *Token) (bool, error) {
	_, affected, err := modifyToken("delete-token", token, nil)
	return affected, err
}

// GetOAuthToken gets the pivotal and necessary secret to interact with the Casdoor server
func GetOAuthToken(code string, state string) (*oauth2.Token, error) {
	config := oauth2.Config{
		ClientID:     authConfig.ClientId,
		ClientSecret: authConfig.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   fmt.Sprintf("%s/api/login/oauth/authorize", authConfig.Endpoint),
			TokenURL:  fmt.Sprintf("%s/api/login/oauth/access_token", authConfig.Endpoint),
			AuthStyle: oauth2.AuthStyleInParams,
		},
		//RedirectURL: redirectUri,
		Scopes: nil,
	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return token, err
	}

	if strings.HasPrefix(token.AccessToken, "error:") {
		return nil, errors.New(strings.TrimLeft(token.AccessToken, "error: "))
	}

	return token, err
}

// RefreshOAuthToken refreshes the OAuth token
func RefreshOAuthToken(refreshToken string) (*oauth2.Token, error) {
	config := oauth2.Config{
		ClientID:     authConfig.ClientId,
		ClientSecret: authConfig.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   fmt.Sprintf("%s/api/login/oauth/authorize", authConfig.Endpoint),
			TokenURL:  fmt.Sprintf("%s/api/login/oauth/refresh_token", authConfig.Endpoint),
			AuthStyle: oauth2.AuthStyleInParams,
		},
		//RedirectURL: redirectUri,
		Scopes: nil,
	}

	token, err := config.TokenSource(context.Background(), &oauth2.Token{RefreshToken: refreshToken}).Token()
	if err != nil {
		return token, err
	}

	if strings.HasPrefix(token.AccessToken, "error:") {
		return nil, errors.New(strings.TrimLeft(token.AccessToken, "error: "))
	}

	return token, err
}

func ValidApiToken(code, secret string) (*User, error) {
	queryMap := map[string]string{
		"code":             code,
		"secret":           secret,
		"application_name": authConfig.ApplicationName,
	}

	url := GetUrl("valid-api-token", queryMap)

	bytes, err := DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var user *User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
