package xstrings
import "testing"
func TestGoletaXstrings004(t *testing.T){if got:=ToKebabCase("HTTPServer");got!="http-server"{t.Fatalf("got=%q",got)}}
