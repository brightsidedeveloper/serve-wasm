package api

import (
	"fmt"
	"io"
	"net/http"
)

func GetWasm() (string, error) {
	resp, err := http.Get("http://localhost:8080/wasm")
	if err != nil {
		return "", fmt.Errorf("failed to get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}
