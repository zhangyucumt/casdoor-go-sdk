package auth

type Block struct {
	*CommonObject

	Type           string                 `xorm:"varchar(32)" json:"type"`
	Comment        string                 `xorm:"mediumtext" json:"comment"`
	UnlockTime     string                 `xorm:"varchar(100)" json:"unlockTime"`
	IP             string                 `xorm:"varchar(100)" json:"ip"`
	NetworkSegment string                 `xorm:"varchar(100)" json:"network_segment"`
	Properties     map[string]interface{} `json:"properties"`
	Status         string                 `xorm:"varchar(100)" json:"status"`
	CreatedTime    string                 `xorm:"varchar(100) created_time created" json:"createdTime"`
}

func GetBlocks() ([]*Block, error) {
	return List[Block]("get-blocks")
}

func GetBlock(name string) (*Block, error) {
	return Get[Block](name, "get-block")
}

func UpdateBlock(block *Block) (bool, error) {
	_, affected, err := Modify[Block]("update-block", *block, nil)
	return affected, err
}

func AddBlock(block *Block) (bool, error) {
	_, affected, err := Modify[Block]("add-block", *block, nil)
	return affected, err
}

func DeleteBlock(block *Block) (bool, error) {
	_, affected, err := Modify[Block]("delete-block", *block, nil)
	return affected, err
}
