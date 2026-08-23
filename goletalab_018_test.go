package xstrings
import "testing"
func TestGoletaXstrings018(t *testing.T){if got:=Translate("hello","aeiou","12345");got!="h2ll4"{t.Fatalf("got=%q",got)}}

func TestGoletaXstrings018AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=Translate("abc","a-c","A-C");got!="ABC"{t.Fatalf("got=%q",got)}
}
