// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yylego/gormcngen

//go:build !gormcngen_generate

// Generated from: gormcnm.gen_test.go:37 -> models_test.TestGenerateColumns
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import (
	"time"

	"github.com/yylego/gormcnm"
	"gorm.io/gorm"
)

func (c *T岗位) Columns() *T岗位Columns {
	return &T岗位Columns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:        gormcnm.Cnm(c.ID, "id"),
		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt: gormcnm.Cnm(c.DeletedAt, "deleted_at"),
		J岗位编号:     gormcnm.Cnm(c.J岗位编号, "job_id"),
		T岗位名称:     gormcnm.Cnm(c.T岗位名称, "title"),
		C公司名称:     gormcnm.Cnm(c.C公司名称, "company"),
		S薪资范围:     gormcnm.Cnm(c.S薪资范围, "salary_range"),
		S薪资下限:     gormcnm.Cnm(c.S薪资下限, "salary_min"),
		S薪资上限:     gormcnm.Cnm(c.S薪资上限, "salary_max"),
		C城市名称:     gormcnm.Cnm(c.C城市名称, "city"),
		L岗位链接:     gormcnm.Cnm(c.L岗位链接, "link"),
		R招聘者:      gormcnm.Cnm(c.R招聘者, "recruiter"),
		E招聘者号:     gormcnm.Cnm(c.E招聘者号, "enc_boss_id"),
		I猎头标记:     gormcnm.Cnm(c.I猎头标记, "is_hunter"),
		S岗位状态:     gormcnm.Cnm(c.S岗位状态, "status"),
		S跳过原因:     gormcnm.Cnm(c.S跳过原因, "skip_reason"),
		M匹配度:      gormcnm.Cnm(c.M匹配度, "match_rate"),
		N备注信息:     gormcnm.Cnm(c.N备注信息, "notes"),
		L最后沟通:     gormcnm.Cnm(c.L最后沟通, "last_comm_at"),
		L最后方向:     gormcnm.Cnm(c.L最后方向, "last_comm_dir"),
		L简历版本:     gormcnm.Cnm(c.L简历版本, "last_resume"),
		D岗位职责:     gormcnm.Cnm(c.D岗位职责, "duties"),
		R岗位要求:     gormcnm.Cnm(c.R岗位要求, "requirements"),
	}
}

type T岗位Columns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	J岗位编号     gormcnm.ColumnName[string]
	T岗位名称     gormcnm.ColumnName[string]
	C公司名称     gormcnm.ColumnName[string]
	S薪资范围     gormcnm.ColumnName[string]
	S薪资下限     gormcnm.ColumnName[int32]
	S薪资上限     gormcnm.ColumnName[int32]
	C城市名称     gormcnm.ColumnName[string]
	L岗位链接     gormcnm.ColumnName[string]
	R招聘者      gormcnm.ColumnName[string]
	E招聘者号     gormcnm.ColumnName[string]
	I猎头标记     gormcnm.ColumnName[bool]
	S岗位状态     gormcnm.ColumnName[E岗位状态]
	S跳过原因     gormcnm.ColumnName[string]
	M匹配度      gormcnm.ColumnName[int32]
	N备注信息     gormcnm.ColumnName[string]
	L最后沟通     gormcnm.ColumnName[int64]
	L最后方向     gormcnm.ColumnName[int32]
	L简历版本     gormcnm.ColumnName[string]
	D岗位职责     gormcnm.ColumnName[string]
	R岗位要求     gormcnm.ColumnName[string]
}

func (c *T需求匹配项) Columns() *T需求匹配项Columns {
	return &T需求匹配项Columns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:        gormcnm.Cnm(c.ID, "id"),
		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt: gormcnm.Cnm(c.DeletedAt, "deleted_at"),
		P岗位主键:     gormcnm.Cnm(c.P岗位主键, "position_id"),
		R岗位要求:     gormcnm.Cnm(c.R岗位要求, "requirement"),
		M匹配状态:     gormcnm.Cnm(c.M匹配状态, "match_status"),
		R简历对应:     gormcnm.Cnm(c.R简历对应, "resume_point"),
		R补充说明:     gormcnm.Cnm(c.R补充说明, "remark"),
		S排序序号:     gormcnm.Cnm(c.S排序序号, "sort_index"),
	}
}

type T需求匹配项Columns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	P岗位主键     gormcnm.ColumnName[uint]
	R岗位要求     gormcnm.ColumnName[string]
	M匹配状态     gormcnm.ColumnName[E匹配状态]
	R简历对应     gormcnm.ColumnName[string]
	R补充说明     gormcnm.ColumnName[string]
	S排序序号     gormcnm.ColumnName[int32]
}

func (c *T沟通记录) Columns() *T沟通记录Columns {
	return &T沟通记录Columns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:        gormcnm.Cnm(c.ID, "id"),
		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt: gormcnm.Cnm(c.DeletedAt, "deleted_at"),
		P岗位主键:     gormcnm.Cnm(c.P岗位主键, "position_id"),
		J岗位编号:     gormcnm.Cnm(c.J岗位编号, "job_id"),
		D消息方向:     gormcnm.Cnm(c.D消息方向, "direction"),
		C消息内容:     gormcnm.Cnm(c.C消息内容, "content"),
		T消息时间:     gormcnm.Cnm(c.T消息时间, "timestamp"),
		B简历消息:     gormcnm.Cnm(c.B简历消息, "is_resume"),
		R简历版本:     gormcnm.Cnm(c.R简历版本, "resume_version"),
	}
}

type T沟通记录Columns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	P岗位主键     gormcnm.ColumnName[uint]
	J岗位编号     gormcnm.ColumnName[string]
	D消息方向     gormcnm.ColumnName[int32]
	C消息内容     gormcnm.ColumnName[string]
	T消息时间     gormcnm.ColumnName[int64]
	B简历消息     gormcnm.ColumnName[bool]
	R简历版本     gormcnm.ColumnName[string]
}

func (c *T黑名单) Columns() *T黑名单Columns {
	return &T黑名单Columns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:        gormcnm.Cnm(c.ID, "id"),
		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt: gormcnm.Cnm(c.DeletedAt, "deleted_at"),
		C公司名称:     gormcnm.Cnm(c.C公司名称, "company"),
		R拉黑原因:     gormcnm.Cnm(c.R拉黑原因, "reason"),
	}
}

type T黑名单Columns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	C公司名称     gormcnm.ColumnName[string]
	R拉黑原因     gormcnm.ColumnName[string]
}
