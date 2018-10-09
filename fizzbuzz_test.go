package main

import "testing"

func TestFizzBuzz_3(t *testing.T) {
	got := fizzbuzz(3)
	want := "Fizz"

	if got != want {
		t.Errorf("fizzbuzz(3) \n got: %v \n want: \n%v", got, want)
	}
}
func TestFizzBuzz_5(t *testing.T) {
	got := fizzbuzz(5)
	want := "Buzz"

	if got != want {
		t.Errorf("fizzbuzz(5) \n got: %v \n want: \n%v", got, want)
	}
}

func TestFizzBuzz_4(t *testing.T) {
	got := fizzbuzz(4)
	want := "4"

	if got != want {
		t.Errorf("fizzbuzz(4) \n got: %v \n want: \n%v", got, want)
	}
}

func TestFizzBuzz_15(t *testing.T) {
	got := fizzbuzz(15)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("fizzbuzz(15) \n got: %v \n want: \n%v", got, want)
	}
}

func TestFizzBuzzTwo_13(t *testing.T) {
	got := fizzbuzztwo(13)
	want := "Fizz"

	if got != want {
		t.Errorf("fizzbuzztwo(13) \n got: %v\n want: %v\n", got, want)
	}
}

func TestFizzBuzzTwo_30(t *testing.T) {
	got := fizzbuzztwo(30)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("fizzbuzztwo(30) \n got: %v\n want: %v\n", got, want)
	}
}

func TestFizzBuzzTwo_57(t *testing.T) {
	got := fizzbuzztwo(57)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("fizzbuzztwo(57) \n got: %v\n want: %v\n", got, want)
	}
}

func TestFizzBuzzTwo_352(t *testing.T) {
	got := fizzbuzztwo(352)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("fizzbuzztwo(352) \n got: %v\n want: %v\n", got, want)
	}
}

func TestFizzBuzzTwo_130(t *testing.T) {
	got := fizzbuzztwo(130)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("fizzbuzztwo(130) \n got: %v\n want: %v\n", got, want)
	}
}

func TestFizzBuzzTwo_224(t *testing.T) {
	got := fizzbuzztwo(224)
	want := "224"

	if got != want {
		t.Errorf("fizzbuzztwo(224) \n got: %v\n want: %v\n", got, want)
	}
}
