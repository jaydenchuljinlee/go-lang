package entity

import "time"

type GySsoUser struct {
	ID         int64     `db:"id"`
	EmailID    string    `db:"email_id"`
	TenantCode string    `db:"tenant_code"`
	DelYn      string    `db:"del_yn"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
