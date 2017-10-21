package domain

type Pack struct {
	packagePath
	packageClass
}

type packageClass string

const (
	corePackage         packageClass = "corePackage"
	internalpackage                  = "projectPackage"
	thirdPartyPackage                = "thirdPartyPackage"
	unknownPackageClass              = "unknownPackageClass"
)

type packagePath string

func (pP packagePath) getPackageClass() (packageClass, error) {
	return corePackage, nil
}

func newPackageFromPath(p string) (*Pack, error) {
	pP := packagePath(p)
	c, err := pP.getPackageClass()
	if err != nil {
		return nil, err
	}

	return &Pack{packagePath: pP, packageClass: c}, nil
}
