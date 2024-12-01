package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/storskegg/poc-kybercast/kybercaster"
)

func main() {
	kc, err := kybercaster.New()
	if err != nil {
		panic(err)
	}
	defer kc.Close()

	// The usual cleanup shit
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	defer func() {
		close(sigChan)
	}()

	go func() {
		<-sigChan

		fmt.Println("Received SIGINT or SIGTERM, shutting down...")
		signal.Stop(sigChan)
		cancel()
	}()

	// Listen for incoming messages
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case data := <-kc.GetChan():
				if data == nil {
					continue
				}

				fmt.Println("=== DATA RECEIVED [RAW] ===")
				spew.Dump(data)
				dec, err := kc.Decrypt(data)
				if err != nil {
					fmt.Printf("Error decrypting data: %s\n", err)
					continue
				}
				fmt.Println("=== DATA RECEIVED [DECRYPTED] ===")
				spew.Dump(dec)
			}
		}
	}()

	t := time.NewTicker(5 * time.Second)
sendLoop:
	for {
		select {
		case <-ctx.Done():
			t.Stop()
			break sendLoop
		case <-t.C:
			_, err := kc.Write([]byte("i am secret data"))
			if err != nil {
				fmt.Printf("!! Error sending data: %s\n", err)
			}
		}
	}

	//privateKey, publicKey, err := kyberk2so.KemKeypair512()
	//if err != nil {
	//	panic(err)
	//}
	//
	//ciphertext, ssA, err := kyberk2so.KemEncrypt512(publicKey)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("=== Ciphertext 1 ===")
	//spew.Dump(ciphertext)
	//
	//ciphertext, ssA, err = kyberk2so.KemEncrypt512(publicKey)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("=== Ciphertext 2 ===")
	//spew.Dump(ciphertext)
	//
	//plaintext := []byte("I am secret data")
	//fmt.Println("=== Plaintext ===")
	//spew.Dump(plaintext)
	//
	//c, err := serpent.NewCipher(ssA[:])
	//if err != nil {
	//	panic(err)
	//}
	//
	//dataEnc := make([]byte, len(plaintext))
	//c.Encrypt(dataEnc, plaintext)
	//fmt.Println("=== Encrypted data ===")
	//spew.Dump(dataEnc)
	//
	//fmt.Println("=== ssa: ===")
	//spew.Dump(ssA)
	//ssB, err := kyberk2so.KemDecrypt512(ciphertext, privateKey)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("=== ssB: ===")
	//spew.Dump(ssB)
	//
	//dataDec := make([]byte, len(dataEnc))
	//d, err := serpent.NewCipher(ssB[:])
	//if err != nil {
	//	panic(err)
	//}
	//d.Decrypt(dataDec, dataEnc)
	//fmt.Println("=== Decrypted data ===")
	//spew.Dump(dataDec)
}
