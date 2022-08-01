package auth

import (
	"encoding/json"
	"fmt"
)

type OperateLogLevel int

const (
	OperateLogLevelDebug OperateLogLevel = iota + 1
	OperateLogLevelInfo
	OperateLogLevelWarn
	OperateLogLevelError
	OperateLogLevelFatal
)

type OperateLog struct {
	Uuid            string          `xorm:"varchar(100) notnull pk" json:"uuid"`
	Application     string          `xorm:"varchar(100) index" json:"application"`
	UserId          string          `xorm:"varchar(100) index" json:"userId"`
	Level           OperateLogLevel `xorm:"tinyint(1) notnull" json:"level"`
	Message         string          `xorm:"mediumtext" json:"message" description:"返回的消息"`
	RequestPath     string          `xorm:"varchar(255)" json:"requestPath" description:"请求的路径"`
	Method          string          `xorm:"varchar(100)" json:"method" description:"请求的方法"`
	Params          string          `xorm:"mediumtext" json:"params" description:"请求的参数"`
	Name            string          `xorm:"varchar(100)" json:"name" description:"操作的名称"`
	HttpCode        int             `xorm:"int" json:"httpCode" description:"请求的http状态码"`
	OperateIp       string          `xorm:"varchar(100) notnull" json:"operateIp"`
	OperateLocation string          `xorm:"varchar(100) notnull" json:"operateLocation"`
	CreatedTime     string          `xorm:"varchar(100)" json:"createdTime"`
}

type OperateLogQueryParams struct {
	PageSize    int             `json:"pageSize"`
	Page        int             `json:"p"`
	Field       string          `json:"field"`
	Value       string          `json:"value"`
	SortField   string          `json:"sortField"`
	SortOrder   string          `json:"sortOrder"`
	User        string          `json:"user"`
	Level       OperateLogLevel `json:"level"`
	RequestPath string          `json:"requestPath"`
	Method      string          `json:"method"`
	Name        string          `json:"name"`
	HttpCode    int             `json:"httpCode"`
	OperateIp   string          `json:"operateIp"`
}

func (p *OperateLogQueryParams) Query() map[string]string {
	pBytes, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(pBytes, &m)
	if err != nil {
		return nil
	}
	ret := make(map[string]string)
	for k, v := range m {
		ret[k] = fmt.Sprintf("%v", v)
	}
	return ret
}

func GetOperateLogs(params OperateLogQueryParams) ([]*OperateLog, error) {
	queryMap := params.Query()

	url := GetUrl("get-operate-logs", queryMap)

	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var operateLogs []*OperateLog
	err = json.Unmarshal(bytes, &operateLogs)
	if err != nil {
		return nil, err
	}
	return operateLogs, nil
}

func GetOperateLog(uuid string) (*OperateLog, error) {
	queryMap := map[string]string{
		"uuid": uuid,
	}

	url := GetUrl("get-operate-log", queryMap)

	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var operateLog *OperateLog
	err = json.Unmarshal(bytes, &operateLog)
	if err != nil {
		return nil, err
	}
	return operateLog, nil
}

func AddOperateLog(operateLog *OperateLog) (bool, error) {
	postBytes, err := json.Marshal(operateLog)
	if err != nil {
		return false, err
	}

	resp, err := doPost("add-operate-log", nil, postBytes, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}
