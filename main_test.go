// You can edit this code!
// Click here and start typing.
package main

import "testing"

func TestMessage(t *testing.T) {
	if Message() == "" {
		t.Fatal("message is empty")
	}
}
