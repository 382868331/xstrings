package xstrings
import "testing"
func TestGoletaXstrings016(t *testing.T){h,m,tail:=LastPartition("plain","/");if h!=""||m!=""||tail!="plain"{t.Fatalf("%q %q %q",h,m,tail)}}
