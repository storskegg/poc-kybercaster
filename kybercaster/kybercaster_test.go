package kybercaster

import (
	"testing"

	kyberk2so "github.com/symbolicsoft/kyber-k2so"

	"golang.org/x/crypto/chacha20poly1305"

	"github.com/awnumar/memguard"
)

func TestKyberCaster_SizeBlockOverhead(t *testing.T) {
	type fields struct {
		PrivateKey *memguard.LockedBuffer
		PublicKey  *memguard.LockedBuffer
		w          chan []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "test",
			fields: fields{},
			want:   chacha20poly1305.Overhead,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KyberCaster{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				w:          tt.fields.w,
			}
			if got := k.SizeBlockOverhead(); got != tt.want {
				t.Errorf("SizeBlockOverhead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyberCaster_SizeCipherText(t *testing.T) {
	type fields struct {
		PrivateKey *memguard.LockedBuffer
		PublicKey  *memguard.LockedBuffer
		w          chan []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "test",
			fields: fields{},
			want:   kyberk2so.Kyber512CTBytes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KyberCaster{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				w:          tt.fields.w,
			}
			if got := k.SizeCipherText(); got != tt.want {
				t.Errorf("SizeCipherText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyberCaster_SizeHeader(t *testing.T) {
	type fields struct {
		PrivateKey *memguard.LockedBuffer
		PublicKey  *memguard.LockedBuffer
		w          chan []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KyberCaster{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				w:          tt.fields.w,
			}
			if got := k.SizeHeader(); got != tt.want {
				t.Errorf("SizeHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyberCaster_SizeNonce(t *testing.T) {
	type fields struct {
		PrivateKey *memguard.LockedBuffer
		PublicKey  *memguard.LockedBuffer
		w          chan []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KyberCaster{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				w:          tt.fields.w,
			}
			if got := k.SizeNonce(); got != tt.want {
				t.Errorf("SizeNonce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyberCaster_SizePrivateKey(t *testing.T) {
	type fields struct {
		PrivateKey *memguard.LockedBuffer
		PublicKey  *memguard.LockedBuffer
		w          chan []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KyberCaster{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				w:          tt.fields.w,
			}
			if got := k.SizePrivateKey(); got != tt.want {
				t.Errorf("SizePrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyberCaster_SizePublicKey(t *testing.T) {
	type fields struct {
		PrivateKey *memguard.LockedBuffer
		PublicKey  *memguard.LockedBuffer
		w          chan []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KyberCaster{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				w:          tt.fields.w,
			}
			if got := k.SizePublicKey(); got != tt.want {
				t.Errorf("SizePublicKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyberCaster_msgIndexCiphertext(t *testing.T) {
	type fields struct {
		PrivateKey *memguard.LockedBuffer
		PublicKey  *memguard.LockedBuffer
		w          chan []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KyberCaster{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				w:          tt.fields.w,
			}
			if got := k.msgIndexCiphertext(); got != tt.want {
				t.Errorf("msgIndexCiphertext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyberCaster_msgIndexData(t *testing.T) {
	type fields struct {
		PrivateKey *memguard.LockedBuffer
		PublicKey  *memguard.LockedBuffer
		w          chan []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KyberCaster{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				w:          tt.fields.w,
			}
			if got := k.msgIndexData(); got != tt.want {
				t.Errorf("msgIndexData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyberCaster_msgIndexNonce(t *testing.T) {
	type fields struct {
		PrivateKey *memguard.LockedBuffer
		PublicKey  *memguard.LockedBuffer
		w          chan []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KyberCaster{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				w:          tt.fields.w,
			}
			if got := k.msgIndexNonce(); got != tt.want {
				t.Errorf("msgIndexNonce() = %v, want %v", got, tt.want)
			}
		})
	}
}
