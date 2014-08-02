
package main 

import "testing"

func TestBirth (t *testing.T) {
    err := Birth("FakeFile.conf.fake")

    if err == nil {
        t.Errorf("Birth() is not handling invalid files correctly")
    }
}
