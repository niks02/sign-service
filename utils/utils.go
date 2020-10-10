package utils

import (
	"encoding/base64"
	"log"
	"math"
)

var txnID uint64 = 0

// EncodeBase64 encodes to base64
func EncodeBase64(message []byte) []byte {
	b := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(b, message)
	return b
}

// DecodeBase64 decode to base64
func DecodeBase64(message []byte) ([]byte, error) {
	var l int
	b := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	l, err := base64.StdEncoding.Decode(b, message)
	if err != nil {
		return nil, err
	}
	return b[:l], nil
}

// GetID return unique id generated for the transaction
func GetID() uint64 {
	if txnID == math.MaxUint64 {
		log.Printf("Resetting transaction ID")
		txnID = 1
	} else {
		txnID++
	}
	return txnID
}
