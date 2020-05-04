package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetMd5(input string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(input))
	if err != nil {
		fmt.Println(err.Error)
	}
	return hex.EncodeToString(hash.Sum(nil))
}
