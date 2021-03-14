package decrypt

import (
	"encoding/base64"
)

func DecodePassword(password string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(password)
	return string(decoded), err
}
