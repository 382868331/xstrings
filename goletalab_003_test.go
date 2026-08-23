package xstrings
import "testing"
func TestGoletaXstrings003(t *testing.T){if got:=ToSnakeCase("HTTPServer");got!="http_server"{t.Fatalf("got=%q",got)}}

func TestGoletaXstrings003AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=ToSnakeCase("FirstName");got!="first_name"{t.Fatalf("got=%q",got)}
}
