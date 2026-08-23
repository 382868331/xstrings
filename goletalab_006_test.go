package xstrings
import "testing"
func TestGoletaXstrings006(t *testing.T){if got:=FirstRuneToUpper("éclair");got!="Éclair"{t.Fatalf("got=%q",got)}}

func TestGoletaXstrings006AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=FirstRuneToUpper("hello");got!="Hello"{t.Fatalf("got=%q",got)}
}
