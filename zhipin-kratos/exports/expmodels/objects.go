package expmodels

import "github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/pkg/models"

// Objects returns all GORM model objects for migration
// 返回所有用于迁移的 GORM 模型对象
func Objects() []any {
	return models.Objects()
}
