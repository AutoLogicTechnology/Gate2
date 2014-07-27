
package main 

import (
    "log"
    "net/http"
    "encoding/json"
    "flag"
    "io/ioutil"
    "fmt"
    "strconv"

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

    if err != nil {
        log.Fatalf("Unable to establish database connection: %s\n", err)
    }

    if g2config.Database.Setup {
        g2config.Database.Connection.AutoMigrate(gate.User{})
        g2config.Database.Connection.AutoMigrate(gate.QRCode{})
        g2config.Database.Connection.AutoMigrate(gate.ScratchCode{})
    }

    goji.Get("/totp/:id/:code", TotpValidateUser)
    goji.Post("/totp/:id", TotpCreateUser)
    goji.Delete("/totp/:id", TotpDeleteUser)
    goji.Put("/totp/:id", TotpUpdateUser)

    // goji.Get("/status", StatusSystem)
    goji.Get("/status/:id", StatusUser)

    goji.Serve()
}

func TotpCreateUser (c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    userid := c.URLParams["id"]
    if !gate.IsValidUserId(userid) {
        i, _ := NewGenericResponse("Failure", "Invalid user ID given")
        http.Error(w, JSONResponse(i), http.StatusBadRequest)
        return 
    }

    var user gate.User

    result := g2config.Database.Connection.Where(&gate.User{UserID: userid}).First(&user)
    if result.Error == nil {
        i, _ := NewGenericResponse("Failure", "That user already exists")
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
        i, _ := NewGenericResponse("Failure", fmt.Sprintf("Unable to add new user: %s (error: %s)", userid, result.Error))
        http.Error(w, JSONResponse(i), http.StatusBadRequest)
        return 
    }

    i, _ := NewTotpResponse("Success", "User created", newuser.QRCode.Base64, newgate.ScratchCodes)

    if g2config.QRCode.WriteToDisk {
        newgate.WritePngToFile(g2config.QRCode.Path + "/" + newgate.UserID + ".png")
    }

    w.Write([]byte(JSONResponse(i)))
}

func TotpValidateUser (c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    userid := c.URLParams["id"]
    totpcode := c.URLParams["code"]

    if !gate.IsValidUserId(userid) {
        i, _ := NewGenericResponse("Failure", "Invalid user ID given")
        http.Error(w, JSONResponse(i), http.StatusBadRequest)
        return
    }

    if !gate.IsValidTOTPCode(totpcode) {
        i, _ := NewGenericResponse("Failure", "Invalid TOTP code given")
        http.Error(w, JSONResponse(i), http.StatusBadRequest)
        return
    }

    var user gate.User
    var codes []gate.ScratchCode

    result := g2config.Database.Connection.Where(&gate.User{UserID: userid}).First(&user)
    if result.Error != nil {
        i, _ := NewGenericResponse("Failure", fmt.Sprintf("Unable to add new user: %s (error: %s)", userid, result.Error))
        http.Error(w, JSONResponse(i), http.StatusNotFound)
        return 
    }

    result = g2config.Database.Connection.Where(&gate.ScratchCode{UserId: user.Id}).Find(&codes)
    if result.Error != nil {
        msg := fmt.Sprintf("Unable to find scratch codes for that userid: %s (error: %s)", userid, result.Error)
        i, _ := NewGenericResponse("Failure", msg)
        http.Error(w, JSONResponse(i), http.StatusNotFound)
        return 
    } 

    g := gate.NewGateWithCustomSecret(user.UserID, user.UserSecret)
    for _, v := range codes {
        g.ScratchCodes = append(g.ScratchCodes, v.Code)

        toint, _ := strconv.Atoi(v.Code)
        g.OTP.ScratchCodes = append(g.OTP.ScratchCodes, toint)
    }

    isvalid, _ := g.CheckCode(totpcode)
    if !isvalid {
        i, _ := NewGenericResponse("Failure", "TOTP code is invalid")

        if len(totpcode) >= 8 {
            i.Message = "TOTP Scratch Code is invalid"
        }

        http.Error(w, JSONResponse(i), http.StatusForbidden)
        return 
    }

    i, _ := NewGenericResponse("Success", "TOTP code is valid")

    if len(totpcode) >= 8 {
        // Scratch code: needs to be deleted
        // after use 
        var sc gate.ScratchCode
        g2config.Database.Connection.Where(&gate.ScratchCode{UserId: user.Id, Code: totpcode}).First(&sc)
        result = g2config.Database.Connection.Delete(&sc)

        i.Message = "TOTP Scratch Code is valid"
    }

    w.Write([]byte(JSONResponse(i)))
}

func TotpDeleteUser (c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    userid := c.URLParams["id"]

    if !gate.IsValidUserId(userid) {
        i, _ := NewGenericResponse("Failure", "Invalid user ID given")
        http.Error(w, JSONResponse(i), http.StatusBadRequest)
        return
    }

    var user gate.User 

    result := g2config.Database.Connection.Where(&gate.User{UserID: userid}).First(&user)
    if result.Error != nil {
        i, _ := NewGenericResponse("Failure", fmt.Sprintf("Unable to find that userid: %s (error: %s)", userid, result.Error))
        http.Error(w, JSONResponse(i), http.StatusNotFound)
        return 
    }

    g2config.Database.Connection.Delete(&user)
    g2config.Database.Connection.Delete(&gate.QRCode{UserId: user.Id})
    g2config.Database.Connection.Delete(&gate.ScratchCode{UserId: user.Id})

    i, _ := NewGenericResponse("Success", "User deleted")
    w.Write([]byte(JSONResponse(i)))
}

func TotpUpdateUser (c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    userid := c.URLParams["id"]
    if !gate.IsValidUserId(userid) {
        i, _ := NewGenericResponse("Failure", "Invalid user ID given")
        http.Error(w, JSONResponse(i), http.StatusBadRequest)
        return
    }

    var user gate.User

    result := g2config.Database.Connection.Where(&gate.User{UserID: userid}).First(&user)
    if result.Error != nil {
        i, _ := NewGenericResponse("Failure", fmt.Sprintf("Unable to find that userid: %s (error: %s)", userid, result.Error))
        http.Error(w, JSONResponse(i), http.StatusNotFound)
        return 
    }

    var codes []gate.ScratchCode

    result = g2config.Database.Connection.Where(&gate.ScratchCode{UserId: user.Id}).Find(&codes)
    if result.Error == nil {
        for _, v := range codes {
            g2config.Database.Connection.Delete(&v)
        }
    }

    user.ScratchCodes = nil
    g2config.Database.Connection.Delete(&user.QRCode)

    i, _ := NewTotpResponse("Success", "User updated", user.QRCode.Base64, []string{""})
    g := gate.NewGateAndQRCode(user.UserID)
    
    user.Generation++
    user.UserSecret = g.UserSecret
    user.QRCode = gate.QRCode{UserId: user.Id, Base64: g.QRCode,}

    for _, v := range g.ScratchCodes {
        g := gate.ScratchCode{UserId: user.Id, Code: v,}
        user.ScratchCodes = append(user.ScratchCodes, g)
        i.ScratchCodes = append(i.ScratchCodes, g.String())
    }

    g2config.Database.Connection.Save(&user)

    w.Write([]byte(JSONResponse(i)))
}

func StatusUser (c web.C, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    userid := c.URLParams["id"]

    var user gate.User
    var codes []gate.ScratchCode

    result := g2config.Database.Connection.Where(&gate.User{UserID: userid}).First(&user)
    if result.Error != nil {
        i, _ := NewGenericResponse("Failure", fmt.Sprintf("Unable to find that userid: %s (error %s)", userid, result.Error))
        http.Error(w, JSONResponse(i), http.StatusNotFound)
        return 
    }

    result = g2config.Database.Connection.Where(&gate.ScratchCode{UserId: user.Id}).Find(&codes)
    if result.Error != nil {
        i, _ := NewGenericResponse("Failure", fmt.Sprintf("Unable to find that userid: %s (error %s)", userid, result.Error))
        http.Error(w, JSONResponse(i), http.StatusNotFound)
        return 
    } 

    i, _ := NewStatusResponse("Success", "User statistics", user.UserID, user.CreatedAt.String(), user.Generation, []string{""})

    for _, v := range codes {
        i.ScratchCodes = append(i.ScratchCodes, v.String())
    }

    w.Write([]byte(JSONResponse(i)))
}
