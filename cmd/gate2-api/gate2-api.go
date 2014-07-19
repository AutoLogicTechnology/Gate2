
package main 

import (
    "log"
    "net/http"
    "encoding/json"
    "flag"
    "io/ioutil"
    "fmt"

    "github.com/AutoLogicTechnology/Gate2/gate"

    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/mattn/go-sqlite3"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

var g2config gate.GateConfiguration

func configuration (filename string) {
    data, err := ioutil.ReadFile(filename)

    if err != nil {
        log.Fatalf("Unable to open the given configuration file: %s\n", filename)
    }

    err = json.Unmarshal(data, &g2config)

    if err != nil {
        log.Fatalf("Unable to read configuration format for: %s: %s\n", filename, err)
    }
}

func main () {
    var err error 

    configfile := flag.String("config", "./gate2.json", "Gate2 configuration file. JSON formatted.")
    flag.Parse()

    configuration(*configfile)

    g2config.Database.Connection, err = gorm.Open(g2config.Database.Engine, g2config.Database.Href)

    if g2config.Database.Purge {
        g2config.Database.Connection.AutoMigrate(gate.User{})
        g2config.Database.Connection.AutoMigrate(gate.QRCode{})
    }

    if err != nil {
        log.Fatalf("Issue opening database: %s: %s\n", g2config.Database.Href, err)
    }

    goji.Get("/totp/:id/:code", TotpValidateUser)
    goji.Post("/totp/:id", TotpCreateUser)
    goji.Delete("/totp/:id", TotpDeleteUser)
    goji.Put("/totp/:id", TotpUpdateUser)

    goji.Serve()
}

func TotpCreateUser (c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    i := TotpCreateUserResponse{}

    userid := c.URLParams["id"]

    if !gate.IsValidUserId(userid) {
        i.Message = "Invalid user ID given"
        http.Error(w, JSONResponse(i), http.StatusBadRequest)
        return 
    }

    var user gate.User 
    result := g2config.Database.Connection.Where(&gate.User{UserID: userid}).First(&user)

    if result.Error == nil {
        i.Message = "That user already exists"
        http.Error(w, JSONResponse(i), http.StatusBadRequest)
        return 
    }

    newgate := gate.NewGateAndQRCode(userid)
    newuser := gate.User{
        Generation: 0,
        UserID: newgate.UserID,
        UserSecret: newgate.UserSecret,
        QRCode: gate.QRCode{
            Base64: newgate.QRCode,
        },
    }

    for _, v := range newgate.ScratchCodes {
        newuser.ScratchCodes = append(newuser.ScratchCodes, gate.ScratchCode{Code: v})
    }

    result = g2config.Database.Connection.Create(&newuser)

    if result.Error != nil {
        i.Message = fmt.Sprintf("Unable to add new user: %s (error: %s)", userid, result.Error)
        http.Error(w, JSONResponse(i), http.StatusBadRequest)
        return 
    }

    i.Message = "User added to the database successfully."
    i.QRCode = newuser.QRCode.Base64

    w.Write([]byte(JSONResponse(i)))
}

func TotpValidateUser (c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    i := TotpValidateUserResponse {
    }

    userid := c.URLParams["id"]
    totpcode := c.URLParams["code"]

    if !gate.IsValidUserId(userid) {
        i.Message = "Invalid user ID given"
        http.Error(w, JSONResponse(i), http.StatusBadRequest)
        return
    }

    if !gate.IsValidTOTPCode(totpcode) {
        i.Message = "Invalid TOTP code given"
        http.Error(w, JSONResponse(i), http.StatusBadRequest)
        return
    }

    var user gate.User
    result := g2config.Database.Connection.Where(&gate.User{UserID: userid}).First(&user)

    if result.Error != nil {
        i.Message = fmt.Sprintf("Unable to find that userid: %s (error %s)", userid, result.Error)
        http.Error(w, JSONResponse(i), http.StatusNotFound)
        return 
    }

    gate := gate.NewGateWithCustomSecret(userid, user.UserSecret)

    isvalid, _ := gate.CheckCode(totpcode)

    if !isvalid {
        i.Message = "TOTP code is invalid"
        http.Error(w, JSONResponse(i), http.StatusForbidden)
        return 
    }

    i.Message = "TOTP code is valid"
    w.Write([]byte(JSONResponse(i)))
}

func TotpDeleteUser (c web.C, w http.ResponseWriter, r *http.Request) {
    i := TotpDeleteUserResponse {Message: "Not Implemented Yet",}
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(JSONResponse(i)))
}

func TotpUpdateUser (c web.C, w http.ResponseWriter, r *http.Request) {
    i := TotpUpdateUserResponse {Message: "Not Implemented Yet",}
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(JSONResponse(i)))
}
