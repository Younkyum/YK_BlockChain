package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"math/big"

	"github.com/Younkyum/nomadcoin/utils"
)

const (
	signature     string = "4d67a6017438f37767d74e7e40a64fbb60d4d9534d2369c76df253d00e392160eb3bca16b72e14aa31e81ede117de21532c682240b3d90dc4b8c2839869b01da%"
	privateKey    string = "307702010104208b121d535d44071e003e7a537c009595909ecfb0909b532cd99266292d4b021ba00a06082a8648ce3d030107a14403420004aeaac5e38cd8947bb53ab76247c6fc926bd00221088e126bbc43ba44433e359dd78f10449bcb88352c4f51793d29285acc00e926569aa6b26b15d87712e5d7c5"
	hashedMessage string = "c33084feaa65adbbbebd0c9bf292a26ffc6dea97b170d88e501ab4865591aafd"
)

func Start() {
	privByte, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	restoredKey, err := x509.ParseECPrivateKey(privByte)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)

	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

}
