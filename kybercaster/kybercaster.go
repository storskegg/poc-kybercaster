package kybercaster

import (
	"net"

	"github.com/awnumar/memguard"
	kyberk2so "github.com/symbolicsoft/kyber-k2so"
	"golang.org/x/crypto/chacha20poly1305"
)

type KyberCaster struct {
	// Kyber
	PrivateKey *memguard.LockedBuffer
	PublicKey  *memguard.LockedBuffer

	w chan []byte

	UDPConn *net.UDPConn
}

func New(address string) (*KyberCaster, error) {
	kc := KyberCaster{}

	kc.w = make(chan []byte)

	kc.PrivateKey = memguard.NewBuffer(kyberk2so.Kyber512SKBytes)
	kc.PublicKey = memguard.NewBuffer(kyberk2so.Kyber512PKBytes)

	privateKey, publicKey, err := kyberk2so.KemKeypair512()
	if err != nil {
		defer kc.Close()

		return nil, err
	}

	kc.PrivateKey.Copy(privateKey[:])
	kc.PublicKey.Copy(publicKey[:])

	kc.PrivateKey.Freeze()
	kc.PublicKey.Freeze()

	// Zero our Arrays
	for i := 0; i < len(privateKey); i++ {
		privateKey[i] = 0
	}

	for i := 0; i < len(publicKey); i++ {
		publicKey[i] = 0
	}

	return &kc, nil
}

func (k *KyberCaster) Close() error {
	k.PrivateKey.Destroy()
	k.PublicKey.Destroy()

	close(k.w)

	return nil
}

func (k *KyberCaster) SizeCipherText() int {
	return kyberk2so.Kyber512CTBytes
}

func (k *KyberCaster) SizePrivateKey() int {
	return kyberk2so.Kyber512SKBytes
}

func (k *KyberCaster) SizePublicKey() int {
	return kyberk2so.Kyber512PKBytes
}

func (k *KyberCaster) SizeBlockOverhead() int {
	return chacha20poly1305.Overhead
}

func (k *KyberCaster) SizeNonce() int {
	return chacha20poly1305.NonceSizeX
}

func (k *KyberCaster) SizeHeader() int {
	return k.SizeCipherText() + k.SizeNonce()
}

func (k *KyberCaster) msgIndexCiphertext() int {
	return 0
}

func (k *KyberCaster) msgIndexNonce() int {
	return k.SizeCipherText()
}

func (k *KyberCaster) msgIndexData() int {
	return k.SizeCipherText() + k.SizeNonce()
}

func (k *KyberCaster) GetChan() <-chan []byte {
	return k.w
}

func (k *KyberCaster) Write(p []byte) (n int, err error) {
	ct, ss, err := kyberk2so.KemEncrypt512([kyberk2so.Kyber512PKBytes]byte(k.PublicKey.Bytes()))
	if err != nil {
		return 0, err
	}

	aead, err := chacha20poly1305.NewX(ss[:])
	if err != nil {
		return 0, err
	}
	nonce := make([]byte, aead.NonceSize())
	enc := aead.Seal(nil, nonce, p, ct[:])

	data := make([]byte, k.SizeHeader()+len(enc))
	copy(data[k.msgIndexCiphertext():], ct[:])
	copy(data[k.msgIndexNonce():], nonce)
	copy(data[k.msgIndexData():], enc[:])

	k.w <- data

	return len(data), nil
}

func (k *KyberCaster) Decrypt(msg []byte) ([]byte, error) {
	ciphertext := [kyberk2so.Kyber512CTBytes]byte(msg[:kyberk2so.Kyber512CTBytes])

	ss, err := kyberk2so.KemDecrypt512(ciphertext, [kyberk2so.Kyber512SKBytes]byte(k.PrivateKey.Bytes()))
	if err != nil {
		return nil, err
	}

	aead, err := chacha20poly1305.NewX(ss[:])
	if err != nil {
		return nil, err
	}

	dec, err := aead.Open(nil, msg[k.msgIndexNonce():k.msgIndexData()], msg[k.msgIndexData():], msg[k.msgIndexCiphertext():k.SizeCipherText()])
	if err != nil {
		return nil, err
	}

	return dec, nil
}
