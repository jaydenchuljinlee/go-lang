package entity

import "time"

type GyCompanyTenant struct {
	ID         int64     `db:"id"`
	TenantCode string    `db:"tenant_code"`
	Workspace  string    `db:"workspace"`
	DelYn      string    `db:"del_yn"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
