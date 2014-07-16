
package otp

import (
    "github.com/AutoLogicTechnology/Gate2/types"
)

func NewTOTPUser (user, secret string) (g *types.Gate, err error) {
    g, _ = types.NewTOTPGate(user, secret)
    return g, nil
}

func AuthenticateTOTPUser (password string) (result bool) {
    g, _ := types.NewTOTPGate(user, nil)
}