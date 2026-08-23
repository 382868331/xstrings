package xstrings
import "testing"
func TestGoletaXstrings017(t *testing.T){if got:=Insert("abcd","X",2);got!="abXcd"{t.Fatalf("got=%q",got)}}

func TestGoletaXstrings017AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=Insert("你好","X",1);got!="你X好"{t.Fatalf("got=%q",got)}
}
