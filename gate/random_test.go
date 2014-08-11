
package gate

import (
    "testing"
    "regexp"
)

type TestTableRow struct {
    ValidRegexp *regexp.Regexp 
    ValidRange string
    Result bool 
}

var RandomStringTable = []TestTableRow {
    TestTableRow{ValidRegexp: regexp.MustCompile("^[abcdefghijk12345]{16}$"), ValidRange: "abcdefghijk12345", Result: true},
    TestTableRow{ValidRegexp: regexp.MustCompile("^[a-z0-9]{16}$"), ValidRange: "abc123___", Result: false},
} 

var ScratchCodeTable = []TestTableRow {
    TestTableRow{ValidRegexp: regexp.MustCompile("^[1-9]{1}[0-9]{7}$"), ValidRange: "", Result: true},
}

var SecretCodeTable = []TestTableRow {
    TestTableRow{ValidRegexp: regexp.MustCompile("^[a-zA-Z0-9!$%^&*()_.,<>?'-]{16}$"), ValidRange: "", Result: true},
}

func TestRandomString(t *testing.T) {
    for _, row := range RandomStringTable {
        randomstr := RandomString(row.ValidRange)

        if row.ValidRegexp.MatchString(randomstr) != row.Result {
            t.Errorf("RandomString() isn't generating acceptable strings: %s", randomstr)
        }
    }
}

func TestNewSecretCode(t *testing.T) {
    for _, row := range SecretCodeTable {
        randomstr := NewSecretCode()

        if row.ValidRegexp.MatchString(randomstr) != row.Result {
            t.Errorf("NewSecretCode() isn't generating acceptable strings: %s", randomstr)
        }
    }
}

func TestScratchCode(t *testing.T) {
    for _, row := range ScratchCodeTable {
        randomstr := NewScratchCode()

        if row.ValidRegexp.MatchString(randomstr) != row.Result {
            t.Errorf("NewScratchCode() isn't generating acceptable strings: %s", randomstr)
        }
    }
}
