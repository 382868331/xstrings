package xstrings
import "testing"
func TestGoletaXstrings019(t *testing.T){if got:=Delete("hello","aeiou");got!="hll"{t.Fatalf("got=%q",got)}}

func TestGoletaXstrings019AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=Delete("abc123","0-9");got!="abc"{t.Fatalf("got=%q",got)}
}
