package fail

import "testing"

func TestFailSample(t *testing.T) {
	t.Run("fail sample test", func(t *testing.T) {
		t.Error("This is success sample error")
	})
}
