package gosigner

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// URLSign genera una firma en string con formato Base64 RawURLEncoding a partir de la infromación
// que se la pasa utilizando SHA-256
func URLSign(data, key []byte) string {

	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
