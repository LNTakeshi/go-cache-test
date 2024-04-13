package fail

import "testing"

func TestFailSample(t *testing.T) {
	t.Run("fail sample test", func(t *testing.T) {
		t.Skip("This is fail sample skip")
		t.Error("This is success sample error")
	})
}
