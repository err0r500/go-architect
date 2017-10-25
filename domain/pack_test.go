package domain

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	currPackage = "github.com/err0r500/go-architect"

	tests := []struct {
		p        string
		expected packageClass
	}{
		{p: "error", expected: corePackage},
		{p: "github.com/err0r500/go-architect", expected: rootPackage},
		{p: "github.com/err0r500/go-architect/interfaces/treeExplorer", expected: projectPackage},
		{p: "github.com/stretchr/testify/assert", expected: thirdPartyPackage},
	}
	for k, tt := range tests {
		returned := packagePath(tt.p).getPackageClass()
		assert.Equal(t, tt.expected, returned, strconv.Itoa(k))
	}
}
