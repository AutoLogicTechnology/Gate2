
package main

type IndexResponse struct {
    HTTPCode int `json:"httpcode"`
    Message string `json:"message"`
}

type TotpIndexResponse struct {
    HTTPCode int `json:"httpcode"`
    Message string `json:"message"`
}

type TotpCreateUserResponse struct {
    HTTPCode int `json:"httpcode"`
    Message string `json:"message"`
}

type TotpValidateUserResponse struct {
    HTTPCode int `json:"httpcode"`
    Message string `json:"message"`
}

type TotpDeleteUserResponse struct {
    HTTPCode int `json:"httpcode"`
    Message string `json:"message"`
}

type TotpUpdateUserResponse struct {
    HTTPCode int `json:"httpcode"`
    Message string `json:"message"`
}