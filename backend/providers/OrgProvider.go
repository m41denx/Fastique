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

func (o *OrgProvider) DeleteLabel(lblId uint) error {
	tx := o.db.Delete(&db.Label{}, lblId)
	return tx.Error
}

func (o *OrgProvider) UpdateLabel(lbl *db.Label) error {
	tx := o.db.Updates(&lbl)
	return tx.Error
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
	tx := o.db.Preload("ServicePoints").Preload("ServicePoints.Labels").Preload("ServicePoints.User").Find(&branches)
	return branches, tx.Error
}

func (o *OrgProvider) CreateSP(branch uint, labels []uint) (uint, error) {
	var lbls []db.Label
	for _, l := range labels {
		lbl := db.Label{}
		lbl.ID = l
		lbls = append(lbls, lbl)
	}
	sp := &db.ServicePoint{
		BranchID: branch,
		Labels:   lbls,
	}
	tx := o.db.Create(sp)
	return sp.ID, tx.Error
}

func (o *OrgProvider) UpdateSP(id uint, labels []uint) error {
	var lbls []db.Label
	for _, l := range labels {
		lbl := db.Label{}
		lbl.ID = l
		lbls = append(lbls, lbl)
	}
	sp := &db.ServicePoint{
		Labels: lbls,
	}
	sp.ID = id
	tx := o.db.Updates(sp)
	return tx.Error
}

func (o *OrgProvider) DeleteSP(id uint) error {
	tx := o.db.Delete(&db.ServicePoint{}, id)
	return tx.Error
}

// endregion

// region Roles

func (o *OrgProvider) CreateRole(name string, privs map[string]interface{}, meta map[string]string) (uint, error) {
	mmeta := utils.CastToIface(meta)
	role := &db.Role{
		Name:       name,
		Meta:       mmeta,
		Privileges: privs,
	}
	tx := o.db.Create(role)
	return role.ID, tx.Error
}

func (o *OrgProvider) ListRoles() ([]db.Role, error) {
	var roles []db.Role
	tx := o.db.Find(&roles)
	return roles, tx.Error
}

// endregion

func (o *OrgProvider) ListUsers() ([]db.User, error) {
	var users []db.User
	tx := o.db.Preload("Role").Find(&users)
	return users, tx.Error
}
