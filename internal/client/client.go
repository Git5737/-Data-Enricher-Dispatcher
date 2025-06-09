package client

import (
	"context"
	"data_enricher_dispatcher/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func FetchUsers(ctx context.Context, url string) ([]model.User, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var users []model.User
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func PostUser(ctx context.Context, url string, user model.User) error {
	body := fmt.Sprintf(`{"name":"%s","email":"%s"}`, user.Name, user.Email)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	for i := 0; i < 3; i++ {
		resp, err = http.DefaultClient.Do(req)
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			resp.Body.Close()
			return nil
		}
		time.Sleep(2 * time.Second)
	}

	return errors.New("failed to POST after 3 retries")
}
