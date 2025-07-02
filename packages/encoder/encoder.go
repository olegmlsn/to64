package encoder

import (
	"encoding/base64"
	"fmt"
	"os"
)

func EncodeFileToBase64(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("unable to read file: %w", err)
	}

	return base64.StdEncoding.EncodeToString(data), nil
}
