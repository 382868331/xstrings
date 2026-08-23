package xstrings
import "testing"
func TestGoletaXstrings013(t *testing.T){h,m,tail:=Partition("a-b-c","-");if h!="a"||m!="-"||tail!="b-c"{t.Fatalf("%q %q %q",h,m,tail)}}
