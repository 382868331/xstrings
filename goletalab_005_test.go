package xstrings
import "testing"
func TestGoletaXstrings005(t *testing.T){if got:=SwapCase("AbC");got!="aBc"{t.Fatalf("got=%q",got)}}

func TestGoletaXstrings005AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=SwapCase("XYZ");got!="xyz"{t.Fatalf("got=%q",got)}
}
