package xstrings
import "testing"
func TestGoletaXstrings007(t *testing.T){if got:=FirstRuneToLower("Éclair");got!="éclair"{t.Fatalf("got=%q",got)}}

func TestGoletaXstrings007AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=FirstRuneToLower("Hello");got!="hello"{t.Fatalf("got=%q",got)}
}
