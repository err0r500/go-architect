package domain

type Pack struct {
	packagePath
	packageClass
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
	corePackage         packageClass = "corePackage"
	internalpackage                  = "projectPackage"
	thirdPartyPackage                = "thirdPartyPackage"
	unknownPackageClass              = "unknownPackageClass"
)

type packagePath string

func (pP packagePath) getPackageClass() packageClass {
	return unknownPackageClass
}
