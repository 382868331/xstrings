package xstrings
import "testing"
func TestGoletaXstrings013(t *testing.T){h,m,tail:=Partition("a-b-c","-");if h!="a"||m!="-"||tail!="b-c"{t.Fatalf("%q %q %q",h,m,tail)}}

func TestGoletaXstrings013AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	h,_,tail:=Partition("one::two::three","::");if h!="one"||tail!="two::three"{t.Fatalf("%q %q",h,tail)}
}
