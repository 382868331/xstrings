package xstrings
import "testing"
func TestGoletaXstrings008(t *testing.T){if got:=Len("你好a");got!=3{t.Fatalf("got=%d",got)}}

func TestGoletaXstrings008AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=Len("éx");got!=2{t.Fatalf("got=%d",got)}
}
