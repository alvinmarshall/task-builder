package domain

type Role struct {
	BaseEntity
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `gorm:"type:uuid" json:"-"`
}

func (Role) TableName() string {
	return "roles"
}

type RoleType int

const (
	USER RoleType = iota
	ADMIN
)

func (r *RoleType) String() string {
	return [...]string{"USER", "ADMIN"}[*r]
}

func (r *RoleType) RoleTypeIndex() int {
	return int(*r)
}
