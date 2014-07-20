
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
    var s string 

    for {
        s = RandomString("0123456789")[:8]

        // s[0] has to be >= 1 for it to be a valid
        // scratch code 
        if !(s[0] == '0') {
            break 
        }
    }

    return s
}

