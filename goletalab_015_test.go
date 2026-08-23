package xstrings
import "testing"
func TestGoletaXstrings015(t *testing.T){h,m,tail:=LastPartition("a-b-c","-");if h!="a-b"||m!="-"||tail!="c"{t.Fatalf("%q %q %q",h,m,tail)}}

func TestGoletaXstrings015AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	h,_,tail:=LastPartition("one::two::three","::");if h!="one::two"||tail!="three"{t.Fatalf("%q %q",h,tail)}
}
