package natsutils

import (
	"crypto/rand"
	"crypto/sha256"
	"io"

	"golang.org/x/crypto/hkdf"
)

func GetKey(masterkey string, nonce [32]byte) ([32]byte, error) {

	masterbytes := ([]byte)(masterkey)

	var key [32]byte
	kdf := hkdf.New(sha256.New, masterbytes, nonce[:], nil)
	if _, err := io.ReadFull(kdf, key[:]); err != nil {
		return [32]byte{}, err
	}

	return key, nil
}

func GetNonce() ([32]byte, error) {

	var nonce [32]byte
	_, err := io.ReadFull(rand.Reader, nonce[:])
	if err != nil {
		return [32]byte{}, err
	}

	return nonce, err
}
