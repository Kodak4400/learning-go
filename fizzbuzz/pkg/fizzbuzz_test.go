package pkg

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	fizzBuzz := FizzBuzz(15)
	if fizzBuzz != "FizzBuzz!" {
		t.Errorf("%v != %v", fizzBuzz, "FizzBuzz!")
	}

	fizz := FizzBuzz(3)
	if fizz != "Fizz!" {
		t.Errorf("%v != %v", fizz, "Fizz!")
	}

	buzz := FizzBuzz(5)
	if buzz != "Buzz!" {
		t.Errorf("%v != %v", buzz, "Buzz!")
	}

	num := FizzBuzz(2)
	if num != "2" {
		t.Errorf("%v != %v", num, "2")
	}

	// これもできるが、、、、
	for _, v := range [4]int{15, 3, 5, 2} {
		switch FizzBuzz(v) {
		case "FizzBuzz!":
			fmt.Println("switch-fizzbuzz", v)
		case "Fizz!":
			fmt.Println("switch-fizz", v)
		case "Buzz!":
			fmt.Println("switch-buzz", v)
		case "2":
			fmt.Println("switch-2", v)
		default:
			t.Errorf("%v != %v", v, "FizzBuzz! or Fizz! or Buzz! or 2")
		}
	}
}
