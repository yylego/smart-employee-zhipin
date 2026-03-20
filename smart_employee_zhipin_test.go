package smart_employee_zhipin_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	smart_employee_zhipin "github.com/yylego/smart-employee-zhipin"
)

func TestSourceRoot(t *testing.T) {
	root := smart_employee_zhipin.SourceRoot()
	t.Log(root)
	require.NotEmpty(t, root)
}
