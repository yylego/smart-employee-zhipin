package models

import "gorm.io/gorm"

// T黑名单 records companies that should be skipped — once rejected, no need to contact again.
//
// T黑名单 记录应跳过的公司——面试被拒后不再联系。
type T黑名单 struct {
	gorm.Model
	C公司名称 string `gorm:"column:company;type:varchar(256);uniqueIndex:idx_blacklist_company;not null"` // 公司名称
	R拉黑原因 string `gorm:"column:reason;type:text"`                               // 拉黑原因
}

func (T黑名单) TableName() string {
	return "blacklist"
}
