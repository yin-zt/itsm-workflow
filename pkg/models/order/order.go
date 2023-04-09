package order

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	GroupName string `gorm:"type:varchar(128);comment:'分组名称'" json:"groupName"`
	Remark    string `gorm:"type:varchar(128);comment:'分组中文说明'" json:"remark"`
	Creator   string `gorm:"type:varchar(20);comment:'创建人'" json:"creator"`
	ParentId  uint   `gorm:"default:0;comment:'父组编号(编号为0时表示根组)'" json:"parentId"`
	Source    string `gorm:"type:varchar(20);comment:'来源：dingTalk、weCom、ldap、platform'" json:"source"`
}
