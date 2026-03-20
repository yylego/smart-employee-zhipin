package zhipin_kratos

import "github.com/yylego/runpath"

// SourceRoot returns the source code root path of zhipin-kratos
func SourceRoot() string {
	return runpath.PARENT.Path()
}
