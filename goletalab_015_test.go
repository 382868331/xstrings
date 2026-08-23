package xstrings
import "testing"
func TestGoletaXstrings015(t *testing.T){h,m,tail:=LastPartition("a-b-c","-");if h!="a-b"||m!="-"||tail!="c"{t.Fatalf("%q %q %q",h,m,tail)}}
