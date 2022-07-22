package auth

import (
	"encoding/json"
	"fmt"
)

type Action struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	Application string `xorm:"varchar(100)" json:"application"`
	CreatedTime string `xorm:"varchar(100) created_time created" json:"createdTime"`

	Label       string `xorm:"varchar(100) notnull" json:"label"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`
	Resource    string `xorm:"mediumtext" json:"resource"`
	Method      string `xorm:"varchar(100)" json:"method"`
}

func GetActions() ([]*Action, error) {
	queryMap := map[string]string{
		"owner":       authConfig.OrganizationName,
		"application": authConfig.ApplicationName,
	}

	url := GetUrl("get-actions", queryMap)

	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var actions []*Action
	err = json.Unmarshal(bytes, &actions)
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func GetAction(name string) (*Action, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", authConfig.OrganizationName, name),
	}

	url := GetUrl("get-action", queryMap)

	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var action *Action
	err = json.Unmarshal(bytes, &action)
	if err != nil {
		return nil, err
	}
	return action, nil
}

func GetActionByUser(username string) (*Action, error) {
	queryMap := map[string]string{
		"owner": authConfig.OrganizationName,
		"user":  username,
	}

	url := GetUrl("get-user-role", queryMap)

	bytes, err := DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var role *Role
	err = json.Unmarshal(bytes, &role)
	if err != nil {
		return nil, err
	}
	return GetAction(role.Name)
}

func UpdateAction(action *Action) (bool, error) {
	_, affected, err := modifyAction("update-action", action, nil)
	return affected, err
}

func AddAction(action *Action) (bool, error) {
	_, affected, err := modifyAction("add-action", action, nil)
	return affected, err
}

func DeleteAction(action *Action) (bool, error) {
	_, affected, err := modifyAction("delete-action", action, nil)
	return affected, err
}
