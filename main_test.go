package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}


func TestMultiply(t *testing.T) {
	total := Multiply(5, 1)
	if total != 5 {
		t.Errorf("Product was incorrect, got: %d, want: %d.", total, 5)
	}
}

func TestReverse(t *testing.T) {
	total := Reverse("hello", "aslan")
	if total != "nalsa olleh" {
		t.Errorf("String was incorrect, got: %d, want: %d.", total, "nalsa olleh")
	}
}