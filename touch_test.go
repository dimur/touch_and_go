package touch

import "testing"

func TestGetMessage(t *testing.T){

    got := GetMessage("Test")
    want := "Test & Go!"

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
