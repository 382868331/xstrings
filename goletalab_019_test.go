package xstrings
import "testing"
func TestGoletaXstrings019(t *testing.T){if got:=Delete("hello","aeiou");got!="hll"{t.Fatalf("got=%q",got)}}
