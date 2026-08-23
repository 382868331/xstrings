package xstrings
import "testing"
func TestGoletaXstrings011(t *testing.T){if got:=Reverse("abc");got!="cba"{t.Fatalf("got=%q",got)}}

func TestGoletaXstrings011AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=Reverse("你好a");got!="a好你"{t.Fatalf("got=%q",got)}
}
