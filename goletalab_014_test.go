package xstrings
import "testing"
func TestGoletaXstrings014(t *testing.T){h,m,tail:=Partition("a--b","--");if h!="a"||m!="--"||tail!="b"{t.Fatalf("%q %q %q",h,m,tail)}}

func TestGoletaXstrings014AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	_,_,tail:=Partition("key=value","=");if tail!="value"{t.Fatalf("tail=%q",tail)}
}
