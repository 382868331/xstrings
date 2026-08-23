package xstrings
import "testing"
func TestGoletaXstrings002(t *testing.T){if got:=ToPascalCase("hello_world");got!="HelloWorld"{t.Fatalf("got=%q",got)}}

func TestGoletaXstrings002AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=ToPascalCase("api_client");got!="ApiClient"{t.Fatalf("got=%q",got)}
}
