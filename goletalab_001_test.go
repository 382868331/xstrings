package xstrings
import "testing"
func TestGoletaXstrings001(t *testing.T){if got:=ToCamelCase("hello_world");got!="helloWorld"{t.Fatalf("got=%q",got)}}
