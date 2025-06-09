package service

import (
	"data_enricher_dispatcher/internal/model"
	"testing"
)

func TestIsBizEmail(t *testing.T) {
	tests := []struct {
		email string
		want  bool
	}{
		{"helloworld@str.biz", true},
		{"some@email.com", false},
		{"my@company.biz", true},
		{"my@.biz", true},
	}
	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			if got := IsBizEmail(tt.email); got != tt.want {
				t.Errorf("IsBizEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterBizUsers(t *testing.T) {
	users := []model.User{
		{Name: "John", Email: "john@company.biz"},
		{Name: "Jane", Email: "jane@gmail.com"},
		{Name: "Bob", Email: "bob@shop.biz"},
	}

	expected := []model.User{
		{Name: "John", Email: "john@company.biz"},
		{Name: "Bob", Email: "bob@shop.biz"},
	}

	result := FilterBizUsers(users)
	if len(result) != len(expected) {
		t.Errorf("Expected %d users, got %d", len(expected), len(result))
	}
}
