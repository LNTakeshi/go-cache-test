package success

import (
	"github.com/samber/lo"
	"testing"
	"time"
)

func TestSample(t *testing.T) {
	t.Run("Sample test", func(t *testing.T) {
		time.Sleep(5 * time.Second)

		s := []string{"a", "b", "c"}
		lo.Map(s, func(v string, k int) string {
			return v
		})
	})
}
