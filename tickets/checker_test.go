package tickets_test

import (
	"testing"

	"github.com/carlriis/Limitless-Lottery/tickets"
)

func TestCheck(t *testing.T) {
	_, got := tickets.Check("####", 1)

	if got == nil {
		t.Error("Did not get an error when checking invalid ID")
	}
}
