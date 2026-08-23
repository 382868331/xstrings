package xstrings
import "testing"
func TestGoletaXstrings012(t *testing.T){if got:=Slice("abcdef",1,4);got!="bcd"{t.Fatalf("got=%q",got)}}

func TestGoletaXstrings012AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=Slice("你好世界",1,3);got!="好世"{t.Fatalf("got=%q",got)}
}
