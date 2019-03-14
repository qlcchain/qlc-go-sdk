package message

import (
	"testing"
)

func TestRandomPoem(t *testing.T) {
	i, poem := RandomPoem()
	t.Log(i, poem.String())
	for idx, v := range poem.Content {
		t.Log(idx, v)
	}
}
