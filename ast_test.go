package parseutil_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zimbabao/go-parse-utils"
)

func TestPackageAST(t *testing.T) {
	pkg, err := parseutil.PackageAST(project)
	require.Nil(t, err)
	require.Equal(t, "parseutil", pkg.Name)
}

func TestPackageTestAST(t *testing.T) {
	pkg, err := parseutil.PackageTestAST(project)
	require.Nil(t, err)
	require.Equal(t, "parseutil_test", pkg.Name)
}
