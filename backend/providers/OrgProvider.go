package providers

import (
	"gorm.io/gorm"
	"main/models/db"
	"main/utils"
)

type OrgProvider struct {
	db *gorm.DB
}

func NewOrgProvider(db *gorm.DB) *OrgProvider {
	return &OrgProvider{
		db: db,
	}
}

func (o *OrgProvider) ExposeGorm() *gorm.DB {
	return o.db
}

// region Labels

func (o *OrgProvider) CreateLabel(name string, tmpl string) (uint, error) {
	label := &db.Label{
		Name:     name,
		Template: tmpl,
	}
	tx := o.db.Create(label)
	return label.ID, tx.Error
}

func (o *OrgProvider) ListLabels() ([]db.Label, error) {
	var labels []db.Label
	tx := o.db.Find(&labels)
	return labels, tx.Error
}

// endregion

//region Branches

func (o *OrgProvider) CreateBranch(name string, meta map[string]string) (uint, error) {
	mmeta := utils.CastToIface(meta)
	branch := &db.Branch{
		Name: name,
		Meta: mmeta,
	}
	tx := o.db.Create(branch)
	return branch.ID, tx.Error
}

func (o *OrgProvider) ListBranches() ([]db.Branch, error) {
	var branches []db.Branch
	tx := o.db.Find(&branches)
	return branches, tx.Error
}

// endregion

// region Roles

func (o *OrgProvider) CreateRole(name string, privs map[string]string, meta map[string]string) (uint, error) {
	mmeta := utils.CastToIface(meta)
	mprivs := utils.CastToIface(privs)
	role := &db.Role{
		Name:       name,
		Meta:       mmeta,
		Privileges: mprivs,
	}
	tx := o.db.Create(role)
	return role.ID, tx.Error
}
