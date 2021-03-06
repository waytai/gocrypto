package hash

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
)

type Digest []byte

var (
	HashAlgo = sha256.New
	HashLen  = 32
)

const ReadSize = 4096

// New computes a new digest computed from the byte slice passed in with the
// algorithm specified by SHA-256.
func New(buf []byte) Digest {
	c := HashAlgo()
	c.Write(buf)
	return c.Sum(nil)
}

// Read computes a new SHA-256 digest from the reader passed in.
func Read(r io.Reader) (h Digest, err error) {
	c := HashAlgo()

	for {
		var n int
		buf := make([]byte, ReadSize)

		n, err = r.Read(buf)
		if err != nil && err != io.EOF {
			return
		}
		c.Write(buf[:n])
		if err == io.EOF {
			err = nil
			break
		}
	}
	h = c.Sum(nil)
	return
}

// The Digest method returns the binary SHA-256 digest.
func (h Digest) Digest() []byte {
	return h
}

// The HexDigest method returns a hexadecimal version of the SHA-256 digest.
func (h Digest) HexDigest() []byte {
	return []byte(fmt.Sprintf("%x", h))
}

// Verify compares the SHA-256 digest to the SHA-256 digest computed from the
// byte slice passed in.
func (h Digest) Verify(buf []byte) bool {
	vHash := New(buf)
	if !bytes.Equal(vHash, h) {
		return false
	}
	return true
}

// VerifyRead compares the SHA-256 digest to the SHA-256 digest computer from
// byte slice passed in.
func (h Digest) VerifyRead(r io.Reader) bool {
	vHash, err := Read(r)
	if err != nil {
		return false
	} else if !bytes.Equal(vHash, h) {
		return false
	}
	return true
}
