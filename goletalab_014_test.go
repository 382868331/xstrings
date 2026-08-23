package xstrings
import "testing"
func TestGoletaXstrings014(t *testing.T){h,m,tail:=Partition("a--b","--");if h!="a"||m!="--"||tail!="b"{t.Fatalf("%q %q %q",h,m,tail)}}
