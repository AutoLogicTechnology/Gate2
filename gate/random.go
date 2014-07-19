
package gate

import (
    "math/rand"
    "time"
)

// This might be a poor means of doing this
// Potentially better(?) ways have been found
// online, but this seems random enough for now 
func RandomString (string_range string) (string) {
    suitable := []byte(string_range)

    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    s := make([]byte, len(suitable))
    p := r.Perm(len(suitable))

    for i, v := range p {
        s[v] = suitable[i]
    }

    return string(s)
}

func NewSecretCode () (string) {
    return RandomString("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!$%^&*()_-.,<>?'")[:16]
}

func NewScratchCode () (string) {
    return RandomString("0123456789")[:8]
}

