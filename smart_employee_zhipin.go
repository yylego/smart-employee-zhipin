package smart_employee_zhipin

import "github.com/yylego/runpath"

// SourceRoot returns the source code root path of smart-employee-zhipin
func SourceRoot() string {
	return runpath.PARENT.Path()
}
