package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/suraj-mahto/fizzbuzz"
)

func TestFizz(t *testing.T) {
	_, ok := fizzbuzz.Fizzbuzz(1)
	if ok{
		t.Logf("The value %v shouldn't have returned true",1)
		t.Fail()
	}

	result, ok := fizzbuzz.Fizzbuzz(3)
	if !ok{
		t.Logf("The value %v shouldn't have returned true",3)
		t.Fail()//other test runs if this fails
	}
	if result != "fizz"{
		t.Log("Result should have been fizz")//only go test -v
		t.Fail() //t,failNow- rest test wont ru
	}

	t.Log("ending fizztest")
}

func ExampleFizzbuzz(){
	result,_:=fizzbuzz.Fizzbuzz(3)
	fmt.Println(result)
	//Output: 3
}