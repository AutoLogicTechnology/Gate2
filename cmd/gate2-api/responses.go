
package main

import (
    "encoding/json"

    
    "github.com/AutoLogicTechnology/Gate2/gate"
)

func JSONResponse (response interface{}) (string) {
    j, _ := json.Marshal(response)
    return string(j)
}

type IndexResponse struct {
    Result string `json:"result"`
    Message string `json:"message"`
}

type TotpIndexResponse struct {
    Result string `json:"result"`
    Message string `json:"message"`
    Gates []*gate.User `json:"gates"`
}

type TotpCreateUserResponse struct {
    Result string `json:"result"`
    Message string `json:"message"`
    QRCode string `json:"qrcode"`
    ScratchCodes []string `json:"scratchcodes"`
}

type TotpValidateUserResponse struct {
    Result string `json:"result"`
    Message string `json:"message"`
}

type TotpDeleteUserResponse struct {
    Result string `json:"result"`
    Message string `json:"message"`
}

type TotpUpdateUserResponse struct {
    Result string `json:"result"`
    Message string `json:"message"`
}

type StatusUserResponse struct {
    Result string `json:"result"`
    Message string `json:"message"`
    UserID string `json:"userid"`
    Created string `json:"created"`
    Generation int64 `json:"generation"`
    ScratchCodes []string `json:"scratchcodes"`
}