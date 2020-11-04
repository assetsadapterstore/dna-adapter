package addrdec

import (
	"fmt"
	"github.com/blocktree/openwallet/openwallet"
	"strings"

	"github.com/blocktree/go-owcdrivers/addressEncoder"
)

var (
	DNAPublicKeyPrefix       = "PUB_"
	DNAPublicKeyK1Prefix     = "PUB_K1_"
	DNAPublicKeyR1Prefix     = "PUB_R1_"
	DNAPublicKeyPrefixCompat = "DNA"

	//DNA stuff
	DNA_mainnetPublic = addressEncoder.AddressType{"dna", addressEncoder.BTCAlphabet, "ripemd160", "", 33, []byte(DNAPublicKeyPrefixCompat), nil}
	// DNA_mainnetPrivateWIF           = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, nil}
	// DNA_mainnetPrivateWIFCompressed = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, []byte{0x01}}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	openwallet.AddressDecoderV2Base
	IsTestNet bool
}

//NewAddressDecoder 地址解析器
func NewAddressDecoderV2() *AddressDecoderV2 {
	decoder := AddressDecoderV2{}
	return &decoder
}

// AddressDecode decode address
func (dec *AddressDecoderV2) AddressDecode(pubKey string, opts ...interface{}) ([]byte, error) {

	var pubKeyMaterial string
	if strings.HasPrefix(pubKey, DNAPublicKeyR1Prefix) {
		pubKeyMaterial = pubKey[len(DNAPublicKeyR1Prefix):] // strip "PUB_R1_"
	} else if strings.HasPrefix(pubKey, DNAPublicKeyK1Prefix) {
		pubKeyMaterial = pubKey[len(DNAPublicKeyK1Prefix):] // strip "PUB_K1_"
	} else if strings.HasPrefix(pubKey, DNAPublicKeyPrefixCompat) { // "DNA"
		pubKeyMaterial = pubKey[len(DNAPublicKeyPrefixCompat):] // strip "DNA"
	} else {
		return nil, fmt.Errorf("public key should start with [%q | %q] (or the old %q)", DNAPublicKeyK1Prefix, DNAPublicKeyR1Prefix, DNAPublicKeyPrefixCompat)
	}

	ret, err := addressEncoder.Base58Decode(pubKeyMaterial, addressEncoder.NewBase58Alphabet(DNA_mainnetPublic.Alphabet))
	if err != nil {
		return nil, addressEncoder.ErrorInvalidAddress
	}
	if addressEncoder.VerifyChecksum(ret, DNA_mainnetPublic.ChecksumType) == false {
		return nil, addressEncoder.ErrorInvalidAddress
	}

	return ret[:len(ret)-4], nil
}

// AddressEncode encode address
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {
	data := addressEncoder.CatData(hash, addressEncoder.CalcChecksum(hash, DNA_mainnetPublic.ChecksumType))
	return string(DNA_mainnetPublic.Prefix) + addressEncoder.EncodeData(data, "base58", DNA_mainnetPublic.Alphabet), nil
}
