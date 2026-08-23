package xstrings
import "testing"
func TestGoletaXstrings004(t *testing.T){if got:=ToKebabCase("HTTPServer");got!="http-server"{t.Fatalf("got=%q",got)}}

func TestGoletaXstrings004AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=ToKebabCase("FirstName");got!="first-name"{t.Fatalf("got=%q",got)}
}
