
package helpers

import (
    "math/rand"
    "time"
)

// This might be a poor means of doing this
// Potentially better(?) ways have been found
// online, but this seems random enough for now 
func RandomString () (string) {
    suitable := []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!$%^&*()_-.,<>?'")

    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    s := make([]byte, len(suitable))
    p := r.Perm(len(suitable))

    for i, v := range p {
        s[v] = suitable[i]
    }

    return string(s)
}