
package main 

import (
    // "log"
    "net/http"
    "encoding/json"

    // "github.com/AutoLogicTechnology/Gate2/gate"
    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

func main () {
    goji.Get("/", Index)

    goji.Get("/totp", http.RedirectHandler("/totp/", 301))
    goji.Get("/totp/", TotpIndex)
    goji.Get("/totp/:id", TotpValidateUser)
    goji.Post("/totp/", TotpCreateUser)
    goji.Delete("/totp/:id", TotpDeleteUser)
    goji.Put("/totp/:id", TotpUpdateUser)

    goji.Serve()
}

func Index (c web.C, w http.ResponseWriter, r *http.Request) {
    i := IndexResponse{
        HTTPCode: 200,
        Message: "Not Implemented Yet",
    }

    j, _ := json.Marshal(i)
    w.Header().Set("Content-Type", "application/json")
    w.Write(j)
}

func TotpIndex (c web.C, w http.ResponseWriter, r *http.Request) {
    i := TotpIndexResponse{
        HTTPCode: 200,
        Message: "Not Implemented Yet",
    }

    j, _ := json.Marshal(i)
    w.Header().Set("Content-Type", "application/json")
    w.Write(j)
}

func TotpCreateUser (c web.C, w http.ResponseWriter, r *http.Request) {
    i := TotpCreateUserResponse {
        HTTPCode: 200,
        Message: "Not Implemented Yet",
    }

    j, _ := json.Marshal(i)
    w.Header().Set("Content-Type", "application/json")
    w.Write(j)
}

func TotpValidateUser (c web.C, w http.ResponseWriter, r *http.Request) {
    i := TotpValidateUserResponse {
        HTTPCode: 200,
        Message: "Not Implemented Yet",
    }

    j, _ := json.Marshal(i)
    w.Header().Set("Content-Type", "application/json")
    w.Write(j)
}

func TotpDeleteUser (c web.C, w http.ResponseWriter, r *http.Request) {
    i := TotpDeleteUserResponse {
        HTTPCode: 200,
        Message: "Not Implemented Yet",
    }

    j, _ := json.Marshal(i)
    w.Header().Set("Content-Type", "application/json")
    w.Write(j)
}

func TotpUpdateUser (c web.C, w http.ResponseWriter, r *http.Request) {
    i := TotpUpdateUserResponse {
        HTTPCode: 200,
        Message: "Not Implemented Yet",
    }

    j, _ := json.Marshal(i)
    w.Header().Set("Content-Type", "application/json")
    w.Write(j)
}
