package string
import "testing"

func Test( t *testing.T ) {
    var Tests = []struct {
        s, want string
    } {
        { "Backward", "drawkcaB"},
        { "Hello, ¡™£¢", "¢£™¡ ,olleH"},
        { "", "" },
    }

    for _, c := range Tests {
        got := Reverse( c.s )
        if got != c.want {
            t.Errorf( "Reverse(%q) == %q, want %q", c.s, got, c.want)
        }
    }
}
