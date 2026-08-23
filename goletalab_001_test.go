package xstrings
import "testing"
func TestGoletaXstrings001(t *testing.T){if got:=ToCamelCase("hello_world");got!="helloWorld"{t.Fatalf("got=%q",got)}}

func TestGoletaXstrings001AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=ToCamelCase("HTTP_server");got!="httpServer"{t.Fatalf("got=%q",got)}
}
