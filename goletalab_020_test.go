package xstrings
import "testing"
func TestGoletaXstrings020(t *testing.T){if got:=Count("hello","aeiou");got!=2{t.Fatalf("got=%d",got)}}

func TestGoletaXstrings020AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=Count("123abc","0-9");got!=3{t.Fatalf("got=%d",got)}
}
