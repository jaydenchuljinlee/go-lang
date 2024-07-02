package service

import (
	"context"
	"fmt"
	"log"
	"toy/domain/dto"
	"toy/redis"
	"toy/repository"
)

func GetTenantUserInfoAndGoogleInfo(ctx context.Context, workspace string, email string) (*dto.GoogleInfo, error) {
	tenantUserInfo, err := repository.GetGyCompanyTenantByEmailAndWorkspace(ctx, workspace, email)

	if err != nil {
		log.Fatalf("Error retrieving tenants: %v", err)
	}

	if tenantUserInfo == nil {
		return nil, nil
	}

	googleInfo, err := redis.GetGoogleInfo(tenantUserInfo.UserID)

	if err != nil {
		return nil, fmt.Errorf("error retrieving Google info: %v", err)
	}

	return googleInfo, nil
}
