
package gate

import "testing"

func TestScratchCodeModel(t *testing.T) {
    sc := ScratchCode{Code: "abc123"}

    if sc.String() != "abc123" {
        t.Error("The ScratchCode{} model doesn't produce the correct String()")
    }
}