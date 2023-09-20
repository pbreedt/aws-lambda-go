package main

import (
	"context"
	"fmt"
	"testing"
)

// run with "go test"  (with -v for t.Log output)

func TestMessage(t *testing.T) {
	event := MyEvent{Name: "Jan", Age: 10}
	resp, _ := HandleLambdaEvent(context.Background(), event)
	if resp.Message != "Jan is 10 years old!" {
		t.Errorf("Msg was incorrect, got: %s, want: %s.", resp.Message, "Jan is 10 years old!")
	}
	t.Log("Success")
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}
