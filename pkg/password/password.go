package password

import (
	"bytes"
	"crypto/rand"
	"github.com/pkg/errors"

	"golang.org/x/crypto/argon2"
)

var DoesntMatchError = errors.New("hash doesn't match")

// HashSalt struct used to store
// generated hash and salt used to
// generate the hash.
type HashSalt struct {
	Hash, Salt []byte
}

func (r HashSalt) GetHash() []byte {
	return r.Hash
}

func (r HashSalt) GetSalt() []byte {
	return r.Salt
}

type Argon2idHash struct {
	// time represents the number of
	// passed over the specified memory.
	time uint32
	// cpu memory to be used.
	memory uint32
	// threads for parallelism aspect
	// of the algorithm.
	threads uint8
	// keyLen of the generate hash key.
	keyLen uint32
	// saltLen the length of the salt used.
	saltLen uint32
}

// NewArgon2idHash constructor function for
// Argon2idHash.
func NewArgon2idHash() *Argon2idHash {
	return &Argon2idHash{
		time:    1,
		saltLen: 32,
		memory:  64 * 1024,
		threads: 32,
		keyLen:  256,
	}
}

// GenerateHash using the password and provided salt.
// If not salt value provided fallback to random value
// generated of a given length.
func (a *Argon2idHash) GenerateHash(password, salt []byte) (*HashSalt, error) {
	var err error
	// If salt is not provided generate a salt of
	// the configured salt length.
	if len(salt) == 0 {
		salt, err = randomSecret(a.saltLen)
	}
	if err != nil {
		return nil, err
	}
	// Generate hash
	hash := argon2.IDKey(password, salt, a.time, a.memory, a.threads, a.keyLen)
	// Return the generated hash and salt used for storage.
	return &HashSalt{Hash: hash, Salt: salt}, nil
}

// Compare generated hash with store hash.
func (a *Argon2idHash) Compare(hash, salt, password []byte) error {
	// Generate hash for comparison.
	hashSalt, err := a.GenerateHash(password, salt)
	if err != nil {
		return err
	}
	// Compare the generated hash with the stored hash.
	// If they don't match return error.
	if !bytes.Equal(hash, hashSalt.Hash) {
		return DoesntMatchError
	}
	return nil
}

func randomSecret(length uint32) ([]byte, error) {
	secret := make([]byte, length)

	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}
