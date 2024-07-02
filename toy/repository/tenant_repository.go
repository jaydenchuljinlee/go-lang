package repository

import (
	"context"
	"database/sql"
	"toy/domain/dto"
)

func GetGyCompanyTenantByEmailAndWorkspace(ctx context.Context, email string, workspace string) (*dto.GyTenantUserInfo, error) {
	query := `
		SELECT c.id AS tenant_id, c.tenant_code, c.workspace, u.id AS user_id, u.email_id AS email
		FROM gy_company_tenant c
		INNER JOIN gy_sso_user u
		ON u.email_id = ? AND u.tenant_code = c.tenant_code AND u.del_yn = 'N'
		WHERE c.workspace = ? AND c.del_yn = 'N'`

	row := DB.QueryRowContext(ctx, query, email, workspace)

	var tenantUserInfo dto.GyTenantUserInfo
	err := row.Scan(&tenantUserInfo.TenantID, &tenantUserInfo.TenantCode, &tenantUserInfo.Workspace, &tenantUserInfo.UserID, &tenantUserInfo.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &tenantUserInfo, nil
}
