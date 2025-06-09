package service

import (
	"context"
	"data_enricher_dispatcher/internal/client"
	"data_enricher_dispatcher/internal/config"
	"data_enricher_dispatcher/internal/model"
	"fmt"
	"log"
	"strings"
)

func IsBizEmail(email string) bool {
	return strings.HasSuffix(email, ".biz")
}

func FilterBizUsers(users []model.User) []model.User {
	var result []model.User
	for _, user := range users {
		if IsBizEmail(user.Email) {
			result = append(result, user)
		}
	}
	return result
}

func ExcludeBizUsers(users []model.User) []model.User {
	var result []model.User
	for _, user := range users {
		if !IsBizEmail(user.Email) {
			result = append(result, user)
		}
	}
	return result
}

func Process(ctx context.Context, cfg config.Config) error {
	users, err := client.FetchUsers(ctx, cfg.ApiA)
	if err != nil {
		return fmt.Errorf("error fetching users: %w", err)
	}

	bizUsers := FilterBizUsers(users)
	skippedUsers := ExcludeBizUsers(users)

	for _, user := range bizUsers {
		if err := client.PostUser(ctx, cfg.ApiB, user); err != nil {
			log.Printf("Failed to POST %s: %v\n", user.Email, err)
		} else {
			log.Printf("Sent: %s\n", user.Email)
		}
	}

	for _, user := range skippedUsers {
		log.Printf("Skipped (not .biz): %s\n", user.Email)
	}

	return nil
}
