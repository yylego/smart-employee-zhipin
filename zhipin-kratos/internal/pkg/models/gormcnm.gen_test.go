package models_test

import (
	"testing"

	"github.com/yylego/gormcngen"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/runpath/runtestpath"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/pkg/models"
)

// Auto generate columns with go generate command
// Support execution via: go generate ./...
// Delete this comment block if auto generation is not needed
//
//go:generate go test -v -run TestGenerateColumns
func TestGenerateColumns(t *testing.T) {
	// Retrieve the absolute path of the source file based on current test file location
	absPath := osmustexist.FILE(runtestpath.SrcPath(t))
	t.Log(absPath)

	// Define data objects used in column generation
	objects := []any{
		&models.T岗位{},
		&models.T匹配项{},
		&models.T沟通记录{},
		&models.T黑名单{},
	}

	// Configure generation options
	options := gormcngen.NewOptions().
		WithColumnClassExportable(true).
		WithColumnsMethodRecvName("c").
		WithColumnsCheckFieldType(true)

	// Create configuration and generate code to target file
	cfg := gormcngen.NewConfigs(objects, options, absPath)
	cfg.Gen()
}
