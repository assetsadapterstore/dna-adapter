package dna

import (
	"github.com/blocktree/bitshares-adapter/encoding"
	"github.com/denkhaus/bitshares/config"
)

//Decrypt calculates a shared secret by the receivers private key
//and the senders public key, then decrypts the given memo message.
func Decrypt(msg []byte, fromPub, toPub string, nonce uint64, wif string) (string, error) {
	ret := config.FindByID(ChainIDDNA)
	if ret == nil {
		config.Add(config.ChainConfig{
			Name:      "DNA",
			CoreAsset: "DNA",
			Prefix:    "DNA",
			ID:        ChainIDDNA,
		})
	}
	config.SetCurrent(ChainIDDNA)

	return encoding.Decrypt(msg, fromPub, toPub, nonce, wif)
}
