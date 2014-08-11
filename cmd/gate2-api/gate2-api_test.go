
package main 

import "testing"

func TestMainApi (t *testing.T) {
    err := MainApi()

    if err == nil {
        t.Error("MainApi() hasn't handled the lack of configuration correctly")
    }
}
