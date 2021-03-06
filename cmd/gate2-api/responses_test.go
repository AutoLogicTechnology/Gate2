
package main 

import (
    "testing"
)

type GenericResponseTableRow struct {
    result string 
    message string

    err bool
    errmsg string
    tostring string 
}

type TotpResponseTableRow struct {
    result string 
    message string 
    qrcode string 
    scratchcodes []string 

    err bool 
    errmsg string  
    tostring string 
}

type StatusResponseTableRow struct {
    result string
    message string 
    userid string 
    created string 
    generation int64 
    scratchcodes []string

    err bool 
    errmsg string 
    tostring string 
}

var GenericResponseTable = []GenericResponseTableRow {
    {"Success", "Successful test 01", false, "", "Success: Successful test 01"},
    {"Success", "Successful test 02", false, "", "Success: Successful test 02"},
    {"Failure", "", true, "Result/message needed", "Success: Successful test 03"},
    {"", "Failed test", true, "Result/message needed", "Success: Successful test 04"},
}

var TotpResponseTable = []TotpResponseTableRow {
    {"Success", "Successful test 01", "qrcode1234", []string{"a","b","c"}, false, "", "Success: Successful test 01 (3)"},
    {"Success", "Successful test 02", "qrcode1234", []string{"a","b","c","d","e","f"}, false, "", "Success: Successful test 02 (6)"},
    {"Failure", "", "qrcode1234", []string{"a","b","c"}, true, "Result/message needed", ""},
    {"", "Failed test", "qrcode1234", []string{"a","b","c"}, true, "Result/message needed", ""},
    {"Failure", "Failed test", "", []string{"a","b","c"}, true, "QRCode needed (base64)", ""},
    {"Failure", "Failed test", "qrcode1234", []string{}, true, "At least one Scratch Code is needed", ""},
}

var StatusResponseTable = []StatusResponseTableRow {
    {"Success", "Successful test", "user01", "today_01", 0, []string{"1","2","3"}, false, "", "Success: Successful test (user01)"},
    {"Success", "", "user02", "today_01", 0, []string{"1","2","3"}, true, "Result/message needed", "Success: Successful test (user02)"},
    {"Success", "Successful test", "", "today_01", 0, []string{"1","2","3"}, true, "User ID needed", "Success: Successful test (user02)"},
    {"Success", "Successful test", "user04", "today_01", 99, []string{"1","2","3"}, false, "", "Success: Successful test (user04)"},
    {"Success", "Successful test", "user05", "today_01", 0, []string{}, true, "At least one Scratch Code is needed", "Success: Successful test (user05)"},
}

func TestNewGenericResponse(t *testing.T) {
    for _, td := range GenericResponseTable {
        actual, err := NewGenericResponse(td.result, td.message)

        if td.err {
            if err == nil {
                t.Errorf("Actual(%s, %s) resulted in no error", td.result, td.message)
            }

            if err.Error() != td.errmsg {
                t.Errorf("Actual(%s, %s) provided the wrong error response: %s", td.result, td.message, td.errmsg)
            }
        } else {
            if actual.Result != td.result {
                t.Errorf("Actual(%s, %s) resulted in the wrong result: %s", actual.Result, actual.Message, td.result)
            }

            if actual.Message != td.message {
                t.Errorf("Actual(%s, %s) resulted in the wrong message: %s", actual.Result, actual.Message, td.message)
            }

            if actual.String() != td.tostring {
                t.Errorf("Actual(%s, %s).String(%s) resulted in wrong String(): %s", 
                    actual.Result, 
                    actual.Message, 
                    actual.String(),
                    td.tostring,
                )
            }
        }
    }
}

func TestNewTotpResponse(t *testing.T) {
    for _, td := range TotpResponseTable {
        actual, err := NewTotpResponse(td.result, td.message, td.qrcode, td.scratchcodes)

        if td.err {
            if err == nil {
                t.Errorf("Actual(%s, %s, %s, %+v) resulted in no error", 
                    td.result, td.message, td.qrcode, td.scratchcodes)
            }

            if err.Error() != td.errmsg {
                t.Errorf("Actual(%+v, %+v, %+v, %+v) provided the wrong error response: %s", 
                    td.result, td.message, td.qrcode, td.scratchcodes, td.errmsg)
            }
        } else {
            if actual.Result != td.result {
                t.Errorf("Actual(%s, %s) resulted in the wrong result: %s", 
                    actual.Result, actual.Message, td.result)
            }

            if actual.Message != td.message {
                t.Errorf("Actual(%s, %s) resulted in the wrong message: %s",
                    actual.Result, actual.Message, td.message)
            }

            if actual.QRCode != td.qrcode {
                t.Errorf("Actual(%s, %s, %s, %+v) resulted in an incorrect QRCode: %s",
                    actual.Result, actual.Message, actual.QRCode, actual.ScratchCodes, td.qrcode)
            }

            if len(actual.ScratchCodes) != len(td.scratchcodes) {
                t.Errorf("Actual(%s, %s, %s, %+v) resulted in an incorrect length of scratch codes: %d",
                    actual.Result, actual.Message, actual.QRCode, actual.ScratchCodes, len(td.scratchcodes))
            }

            if actual.String() != td.tostring {
                t.Errorf("Actual(%s, %s, %s, %+v) resulted in an incorrect .String(): %s",
                    actual.Result, actual.Message, actual.QRCode, actual.ScratchCodes, len(td.scratchcodes), td.tostring)
            }
        }
    }
}

func TestNewStatusResponse(t *testing.T) {
    for _, td := range StatusResponseTable {
        actual, err := NewStatusResponse(td.result, td.message, td.userid, td.created, td.generation, td.scratchcodes)

        if td.err {
            if err == nil {
                t.Errorf("Actual(%s, %s, %s, %s, %d, %+v) resulted in no error",
                    td.result, 
                    td.message, 
                    td.userid, 
                    td.created, 
                    td.generation, 
                    td.scratchcodes,
                )
            }

            if err.Error() != td.errmsg {
                t.Errorf("Actual(%s, %s, %s, %s, %d, %+v) provided the wrong error response: %s",
                    td.result,
                    td.message,
                    td.userid,
                    td.created,
                    td.generation,
                    td.scratchcodes,
                    td.errmsg,
                )
            }
        } else {
            if actual.Result != td.result {
                t.Errorf("Actual(%s, %s, %s, %s, %d, %+v) resulted in the wrong result: %s",
                    actual.Result,
                    actual.Message,
                    actual.UserID,
                    actual.Created,
                    actual.Generation,
                    actual.ScratchCodes,
                    td.result,
                )
            }

            if actual.Message != td.message {
                t.Errorf("Actual(%s, %s, %s, %s, %d, %+v) resulted in the wrong message: %s",
                    actual.Result,
                    actual.Message,
                    actual.UserID,
                    actual.Created,
                    actual.Generation,
                    actual.ScratchCodes,
                    td.message,
                )
            }

            if actual.UserID != td.userid {
                t.Errorf("Actual(%s, %s, %s, %s, %d, %+v) resulted in an incorrect user id: %s",
                    actual.Result,
                    actual.Message,
                    actual.UserID,
                    actual.Created,
                    actual.Generation,
                    actual.ScratchCodes,
                    td.userid,
                )
            }   

            if actual.Created != td.created {
                t.Errorf("Actual(%s, %s, %s, %s, %d, %+v) resulted in an incorrect created at: %s",
                    actual.Result,
                    actual.Message,
                    actual.UserID,
                    actual.Created,
                    actual.Generation,
                    actual.ScratchCodes,
                    td.created,
                )
            }

            if actual.Generation != td.generation {
                t.Errorf("Actual(%s, %s, %s, %s, %d, %+v) resulted in an incorrect generation: %d",
                    actual.Result,
                    actual.Message,
                    actual.UserID,
                    actual.Created,
                    actual.Generation,
                    actual.ScratchCodes,
                    td.generation,
                )
            }

            if len(actual.ScratchCodes) != len(td.scratchcodes) {
                t.Errorf("Actual(%s, %s, %s, %s, %d, %+v) resulted in an incorrect length of scratch codes: %d",
                    actual.Result,
                    actual.Message,
                    actual.UserID,
                    actual.Created,
                    actual.Generation,
                    actual.ScratchCodes,
                    len(td.scratchcodes),
                )
            }

            if actual.String() != td.tostring {
                t.Errorf("Actual(%s, %s, %s, %s, %d, %+v) resulted in an incorrect String(): %s",
                    actual.Result,
                    actual.Message,
                    actual.UserID,
                    actual.Created,
                    actual.Generation,
                    actual.ScratchCodes,
                    td.tostring,
                )
            }
        }
    }
}