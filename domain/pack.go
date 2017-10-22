package domain

import (
	"go/build"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var currPackage string

func init() {
	setCurrPackageImportPath()
}

type Pack struct {
	packagePath
	packageClass
}

func (p Pack) String() string {
	return string(p.packagePath) + " (" + string(p.packageClass) + ")"
}

func NewPackFromPath(p string) *Pack {
	pP := packagePath(p)
	return &Pack{
		packagePath:  pP,
		packageClass: pP.getPackageClass(),
	}
}

type packageClass string

const (
	corePackage       packageClass = "corePackage"
	internalpackage                = "projectPackage"
	thirdPartyPackage              = "thirdPartyPackage"
)

type packagePath string

func (pP packagePath) getPackageClass() packageClass {
	var internal = regexp.MustCompile(currPackage + `.*`)
	if internal.MatchString(string(pP)) {
		return internalpackage
	}

	var core = regexp.MustCompile(`"[a-z]*[^/]"`) // pas très classe et souvent plus compliqué que ça
	if core.MatchString(string(`"` + pP + `"`)) {
		return corePackage
	}

	return thirdPartyPackage
}

func setCurrPackageImportPath() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	currPackage = strings.Replace(filepath.Dir(ex), build.Default.GOPATH+"/src/", "", -1)
}
