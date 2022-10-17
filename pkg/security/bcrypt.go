// For simplicity i don't write unit test
package security

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
	Cost int
}

func (b *Bcrypt) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), b.Cost)
	return string(bytes), err
}

func (b *Bcrypt) Verify(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
