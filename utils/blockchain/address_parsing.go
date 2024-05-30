package utils

import (
	"log"
	"github.com/KimDaeikk/filmountain-sdk/config"
)

// hex to []byte
func StringToByteEthAddr(hexAddr string) [20]byte {
	ethAddr, err := ethtypes.ParseEthAddress(hexAddr)
	if err != nil {
		log.Fatal(err)
	}
	return ethAddr
}

func (byteEthAddr *[20]byte) ByteEthAddrToFilAddr() {
	byteEthAddr.ToFilecoinAddress()
	if err != nil {
		logFatal(err)
	}
}

