package keys

import (
	"encoding/base32"
	"encoding/hex"
	"log"
	"math/big"
)

// GetAccountIDFromAccessKey
// refs https://trufflesecurity.com/blog/canaries
func GetAccountIDFromAccessKey(accessKey string) (int64, error) {
	prefixRemoved := accessKey[4:]
	decoded, err := base32.StdEncoding.DecodeString(prefixRemoved)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	idBytes := decoded[0:6]

	z := new(big.Int).SetBytes(idBytes)
	mask, _ := hex.DecodeString("7fffffffff80")
	maskInt := new(big.Int).SetBytes(mask)

	e := new(big.Int).Rsh(new(big.Int).And(z, maskInt), 7)
	return e.Int64(), nil
}
