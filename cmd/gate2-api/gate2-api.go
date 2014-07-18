
package main 

import (
    // "log"
    "net/http"
    "encoding/json"
    "fmt"

    "github.com/AutoLogicTechnology/Gate2/gate"
    "github.com/AutoLogicTechnology/Gate2/database"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

var gates []*TotpIndexResponseGate

func main () {
    goji.Get("/", Index)

    goji.Get("/totp", http.RedirectHandler("/totp/", 301))
    goji.Get("/totp/", TotpIndex)
    goji.Get("/totp/:id/:code", TotpValidateUser)
    goji.Post("/totp/:id", TotpCreateUser)
    goji.Delete("/totp/:id", TotpDeleteUser)
    goji.Put("/totp/:id", TotpUpdateUser)

    web.C.Env["GATE_DATABASE"] = &database.GateSQLiteDatabase{
        
    }

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
        Message: "Current Gates",
    }

    w.Header().Set("Content-Type", "application/json")

    if len(gates) >= 1 {
        for _, gate := range gates {
            s := fmt.Sprintf("%s:%s", gate.UserID, gate.UserSecret)
            i.Gates = append(i.Gates, &s)
        }
    }

    j, _ := json.Marshal(i)
    w.Write(j)
}

func TotpCreateUser (c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    i := TotpCreateUserResponse {
        HTTPCode: 200,
    }

    userid := c.URLParams["id"]

    if !gate.IsValidUserId(userid) {
        http.Error(w, "Invalid user ID given", http.StatusBadRequest)
        return 
    }

    if haveuser(userid) {
        http.Error(w, "User exists", http.StatusBadRequest)
        return 
    }

    newgate := gate.NewGateAndQRCode(userid)
    gates = append(gates, &TotpIndexResponseGate{
        UserID: newgate.UserID,
        UserSecret: newgate.UserSecret,
    })

    i.Message = "User added to the database successfully."

    j, _ := json.Marshal(i)
    
    w.Write(j)
}

func TotpValidateUser (c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    i := TotpValidateUserResponse {
        HTTPCode: 200,
    }

    userid := c.URLParams["id"]
    totpcode := c.URLParams["code"]

    if !gate.IsValidUserId(userid) {
        http.Error(w, "Invalid user ID given", http.StatusBadRequest)
        return
    }

    if !gate.IsValidTOTPCode(totpcode) {
        http.Error(w, "Invalid TOTP code given", http.StatusBadRequest)
        return
    }



    j, _ := json.Marshal(i)
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
