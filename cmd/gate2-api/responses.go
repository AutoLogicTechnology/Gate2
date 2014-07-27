
package main

import (
    "encoding/json"
    "errors"
    "time"
    "fmt"
)

type GenericResponse struct {
    Result string `json:"result"`
    Message string `json:"message"`
}

type TotpResponse struct {
    Result string `json:"result"`
    Message string `json:"message"`
    QRCode string `json:"qrcode"`
    ScratchCodes []string `json:"scratchcodes"`
}

type StatusResponse struct {
    Result string `json:"result"`
    Message string `json:"message"`
    UserID string `json:"userid"`
    Created string `json:"created"`
    Generation int64 `json:"generation"`
    ScratchCodes []string `json:"scratchcodes"`
}

func JSONResponse (response interface{}) (string) {
    j, _ := json.Marshal(response)
    return string(j)
}

func (r GenericResponse) String() (string) {
    return fmt.Sprintf("%s: %s", r.Result, r.Message)
}

func (r TotpResponse) String() (string) {
    return fmt.Sprintf("%s: %s (%d)", r.Result, r.Message, len(r.ScratchCodes))
}

func (r StatusResponse) String() (string) {
    return fmt.Sprintf("%s: %s (%s)", r.Result, r.Message, r.UserID)
}

func NewGenericResponse(result, message string) (*GenericResponse, error) {
    x := &GenericResponse{}

    if result == "" || message == "" {
        return nil, errors.New("Result/message needed")
    } 

    x.Result, x.Message = result, message
    return x, nil
}

func NewTotpResponse(result, message, qrcode string, scratchcodes []string) (*TotpResponse, error) {
    x := &TotpResponse{}

    if result == "" || message == "" {
        return nil, errors.New("Result/message needed")
    }

    if qrcode == "" {
        return nil, errors.New("QRCode needed (base64)")
    }

    if len(scratchcodes) <= 0 {
        return nil, errors.New("At least one Scratch Code is needed")
    }

    x.Result, x.Message, x.QRCode, x.ScratchCodes = result, message, qrcode, scratchcodes
    return x, nil 
}

func NewStatusResponse(result, message, userid, created string, generation int64, scratchcodes []string) (*StatusResponse, error) {
    x := &StatusResponse{}

    if result == "" || message == "" {
        return nil, errors.New("Result/message needed")
    }

    x.Result = result 
    x.Message = message 

    if userid == "" {
        return nil, errors.New("User ID needed")   
    } else {
        x.UserID = userid 
    }

    if created == "" {
        x.Created = time.Now().String()
    } else {
        x.Created = created 
    }

    if generation <= 0 {
        x.Generation = 0
    } else {
        x.Generation = generation
    }

    if len(scratchcodes) <= 0 {
        return nil, errors.New("At least one Scratch Code is needed")
    } else {
        x.ScratchCodes = scratchcodes
    }
    
    return x, nil
}
