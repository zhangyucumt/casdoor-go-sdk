package auth

import (
	"encoding/json"
)

type Provider struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	DisplayName       string `xorm:"varchar(100)" json:"displayName"`
	Category          string `xorm:"varchar(100)" json:"category"`
	Type              string `xorm:"varchar(100)" json:"type"`
	SubType           string `xorm:"varchar(100)" json:"subType"`
	Method            string `xorm:"varchar(100)" json:"method"`
	ClientId          string `xorm:"varchar(100)" json:"clientId"`
	ClientSecret      string `xorm:"varchar(2000)" json:"clientSecret"`
	ClientId2         string `xorm:"varchar(100)" json:"clientId2"`
	ClientSecret2     string `xorm:"varchar(100)" json:"clientSecret2"`
	Cert              string `xorm:"varchar(100)" json:"cert"`
	CustomAuthUrl     string `xorm:"varchar(200)" json:"customAuthUrl"`
	CustomScope       string `xorm:"varchar(200)" json:"customScope"`
	CustomTokenUrl    string `xorm:"varchar(200)" json:"customTokenUrl"`
	CustomUserInfoUrl string `xorm:"varchar(200)" json:"customUserInfoUrl"`
	CustomLogo        string `xorm:"varchar(200)" json:"customLogo"`

	Host    string `xorm:"varchar(100)" json:"host"`
	Port    int    `json:"port"`
	Title   string `xorm:"varchar(100)" json:"title"`
	Content string `xorm:"varchar(1000)" json:"content"`

	RegionId     string `xorm:"varchar(100)" json:"regionId"`
	SignName     string `xorm:"varchar(100)" json:"signName"`
	TemplateCode string `xorm:"varchar(100)" json:"templateCode"`
	AppId        string `xorm:"varchar(100)" json:"appId"`

	Endpoint         string `xorm:"varchar(1000)" json:"endpoint"`
	IntranetEndpoint string `xorm:"varchar(100)" json:"intranetEndpoint"`
	Domain           string `xorm:"varchar(100)" json:"domain"`
	Bucket           string `xorm:"varchar(100)" json:"bucket"`

	Metadata               string `xorm:"mediumtext" json:"metadata"`
	IdP                    string `xorm:"mediumtext" json:"idP"`
	IssuerUrl              string `xorm:"varchar(100)" json:"issuerUrl"`
	EnableSignAuthnRequest bool   `json:"enableSignAuthnRequest"`

	ProviderUrl string `xorm:"varchar(200)" json:"providerUrl"`
}

type ProviderItem struct {
	Name      string    `json:"name"`
	CanSignUp bool      `json:"canSignUp"`
	CanSignIn bool      `json:"canSignIn"`
	CanUnlink bool      `json:"canUnlink"`
	Prompted  bool      `json:"prompted"`
	AlertType string    `json:"alertType"`
	Provider  *Provider `json:"provider"`
}

func GetProviders(owner string) ([]Provider, error) {
	queryMap := map[string]string{
		"owner": owner,
	}
	url := GetUrl("get-providers", queryMap)
	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}
	var providers []Provider
	err = json.Unmarshal(bytes, &providers)
	if err != nil {
		return nil, err
	}
	return providers, nil
}
