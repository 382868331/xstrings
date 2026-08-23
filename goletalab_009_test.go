package xstrings
import "testing"
func TestGoletaXstrings009(t *testing.T){if got:=WordCount("hello world");got!=2{t.Fatalf("got=%d",got)}}

func TestGoletaXstrings009AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=WordCount("");got!=0{t.Fatalf("got=%d",got)}
}
