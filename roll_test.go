package odds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type rollTest struct {
	diff      int
	dice      []int
	botch     bool
	successes int
}

func TestRoll(t *testing.T) {
	assert := assert.New(t)

	tests := []rollTest{
		rollTest{5, []int{1, 1}, true, 0},
		rollTest{5, []int{1, 2, 3, 4}, true, 0},
		rollTest{5, []int{1, 5}, false, 0},
		rollTest{5, []int{1, 1, 5}, false, 0},
		rollTest{5, []int{2, 2}, false, 0},
		rollTest{5, []int{2, 5}, false, 1},
		rollTest{5, []int{5, 5}, false, 2},
	}
	for _, test := range tests {
		successes, botch := Interpret(test.diff, test.dice)
		assert.Equal(test.botch, botch, "botch")
		assert.Equal(test.successes, successes, "successes")
	}
}
