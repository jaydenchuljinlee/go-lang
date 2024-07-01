package dto

type GyTenantUserInfo struct {
	TenantID   int64  `db:"tenant_id"`
	TenantCode string `db:"tenant_code"`
	Workspace  string `db:"workspace"`
	UserID     int64  `db:"user_id"`
	Email      string `db:"email"`
}
