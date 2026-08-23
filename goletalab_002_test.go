package xstrings
import "testing"
func TestGoletaXstrings002(t *testing.T){if got:=ToPascalCase("hello_world");got!="HelloWorld"{t.Fatalf("got=%q",got)}}
