package auth

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CommonObject struct {
	Owner string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name  string `xorm:"varchar(100) notnull pk" json:"name"`
}

func (c *CommonObject) GetOwner() string {
	return c.Owner
}

func (c *CommonObject) GetName() string {
	return c.Name
}

func (c *CommonObject) SetOwner(owner string) {
	c.Owner = owner
}

type obj interface {
	GetOwner() string
	GetName() string
	SetOwner(owner string)
}

func List[T any](uri string) ([]*T, error) {
	queryMap := map[string]string{
		"owner": authConfig.OrganizationName,
	}

	url := GetUrl(uri, queryMap)

	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var data []*T
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Get[T any](name, uri string) (*T, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", authConfig.OrganizationName, name),
	}

	url := GetUrl(uri, queryMap)

	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var data *T
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Modify is an encapsulation of user CUD(Create, Update, Delete) operations.
func Modify[T obj, PT interface {
	GetOwner() string
	GetName() string
	SetOwner(string)
	*T
}](url string, t T, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", t.GetOwner(), t.GetName()),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	t.SetOwner(authConfig.OrganizationName)
	postBytes, err := json.Marshal(t)
	if err != nil {
		return nil, false, err
	}

	resp, err := doPost(url, queryMap, postBytes, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}
