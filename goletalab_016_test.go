package xstrings
import "testing"
func TestGoletaXstrings016(t *testing.T){h,m,tail:=LastPartition("plain","/");if h!=""||m!=""||tail!="plain"{t.Fatalf("%q %q %q",h,m,tail)}}

func TestGoletaXstrings016AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	h,_,tail:=LastPartition("abc","x");if h!=""||tail!="abc"{t.Fatalf("%q %q",h,tail)}
}
