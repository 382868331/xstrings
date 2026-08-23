package xstrings
import "testing"
func TestGoletaXstrings010(t *testing.T){if got:=RuneWidth(rune(10));got!=0{t.Fatalf("got=%d",got)}}

func TestGoletaXstrings010AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=RuneWidth(rune(9));got!=0{t.Fatalf("got=%d",got)}
}
