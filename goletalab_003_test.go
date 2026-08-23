package xstrings
import "testing"
func TestGoletaXstrings003(t *testing.T){if got:=ToSnakeCase("HTTPServer");got!="http_server"{t.Fatalf("got=%q",got)}}
