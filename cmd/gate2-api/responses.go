
package main

import (
    "encoding/json"

    
    "github.com/AutoLogicTechnology/Gate2/gate"
)

type IndexResponse struct {
    Message string `json:"message"`
}

type TotpIndexResponse struct {
    Message string `json:"message"`

    Gates []*gate.User `json:"gates"`
}

type TotpCreateUserResponse struct {
    Message string `json:"message"`
    QRCode string `json:"qrcode"`
    ScratchCodes []string `json:"scratchcodes"`
}

type TotpValidateUserResponse struct {
    Message string `json:"message"`
}

type TotpDeleteUserResponse struct {
    Message string `json:"message"`
}

type TotpUpdateUserResponse struct {
    Message string `json:"message"`
}

func JSONResponse (response interface{}) (string) {
    j, _ := json.Marshal(response)
    return string(j)
}