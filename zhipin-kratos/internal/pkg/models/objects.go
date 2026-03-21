package models

// Objects returns all GORM model objects for migration
// 返回所有用于迁移的 GORM 模型对象
func Objects() []any {
	return []any{
		&T岗位{},
		&T需求匹配项{},
		&T沟通记录{},
		&T黑名单{},
	}
}
