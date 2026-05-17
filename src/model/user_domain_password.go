package model

import (
	"crypto/sha256"
	"encoding/hex"
)

 func (ud *userDomain) EncryptPassword() {
	hash := sha256.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}