package hello
import "testing"

func TestHello(t *testing.T){
	expected:="Hello Go!"
	actual := hello()
	if actual != expected{
		t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
	}
}
func TestHella(t *testing.T){
	expected:="Hello Go"
	actual := hella()
	if actual != expected{
		t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
	}
}

