package values

import (
	"testing"
	"time"

	"github.com/vvanpo/vind/types"
)

func TestTimeEpoch(t *testing.T) {
	sec := types.Time(0)
	out := TimeFormat{}.Format(sec)

	if out != "1970-01-01 00:00:00" {
		t.Errorf("%s does not match epoch", out)
	}

	out = TimeFormat{Location: time.FixedZone("foo", -3600*3)}.Format(sec)
	if out != "1969-12-31 21:00:00" {
		t.Errorf("%s does not match epoch offset -0300", out)
	}
}
