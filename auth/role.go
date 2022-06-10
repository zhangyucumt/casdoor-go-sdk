package auth

import (
	"encoding/json"
	"fmt"
)

type Role struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`

	Users     []string `xorm:"mediumtext" json:"users"`
	Roles     []string `xorm:"mediumtext" json:"roles"`
	IsEnabled bool     `json:"isEnabled"`
}

func GetRoles() ([]*Role, error) {
	queryMap := map[string]string{
		"owner": authConfig.OrganizationName,
	}

	url := GetUrl("get-roles", queryMap)

	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var roles []*Role
	err = json.Unmarshal(bytes, &roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func GetRole(name string) (*Role, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", authConfig.OrganizationName, name),
	}

	url := GetUrl("get-role", queryMap)

	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var role *Role
	err = json.Unmarshal(bytes, &role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func GetRoleByUser(username string) (*Role, error) {
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
	return role, nil
}

func UpdateRole(role *Role) (bool, error) {
	_, affected, err := modifyRole("update-role", role, nil)
	return affected, err
}

func AddRole(role *Role) (bool, error) {
	_, affected, err := modifyRole("add-role", role, nil)
	return affected, err
}

func DeleteRole(role *Role) (bool, error) {
	_, affected, err := modifyRole("delete-role", role, nil)
	return affected, err
}
