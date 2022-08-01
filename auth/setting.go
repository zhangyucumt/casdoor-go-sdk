package auth

type Setting struct {
	*CommonObject

	Value       string `xorm:"notnull" json:"value"`
	Description string `xorm:"varchar(255) notnull" json:"description"`

	CreatedTime string `xorm:"varchar(100) created_time created" json:"createdTime"`
}

func GetSettings() ([]*Setting, error) {
	return List[Setting]("get-settings")
}

func GetSetting(name string) (*Setting, error) {
	return Get[Setting](name, "get-setting")
}

func UpdateSetting(s *Setting) (bool, error) {
	_, affected, err := Modify[Setting]("update-setting", *s, nil)
	return affected, err
}

func AddSetting(s *Setting) (bool, error) {
	_, affected, err := Modify[Setting]("add-setting", *s, nil)
	return affected, err
}

func DeleteSetting(s *Setting) (bool, error) {
	_, affected, err := Modify[Setting]("delete-setting", *s, nil)
	return affected, err
}
