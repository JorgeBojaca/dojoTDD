package calculator
import "testing"

func TestAdd_1_1(t *testing.T){
	expected:=2
	actual:=add(1,1)
	if expected != actual{
		t.Errorf("Test failed, expected: '%d', got: '%d'",expected,actual)

	}
}
func TestAdd_3_2(t *testing.T){
	expected:=5
	actual:=add(3,2)
	if expected != actual{
		t.Errorf("Test failed, expected: '%d', got: '%d'",expected,actual)

	}
}
func TestSub_3_2(t *testing.T){
	expected:=1
	actual:=sub(3,2)
	if expected != actual{
		t.Errorf("Test failed, expected: '%d', got: '%d'",expected,actual)

	}
}
func TestSub_1_5(t *testing.T){
	expected:=-4
	actual:=sub(1,5)
	if expected != actual{
		t.Errorf("Test failed, expected: '%d', got: '%d'",expected,actual)

	}
}