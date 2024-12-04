package encoding

import (
	"fmt"
	"math"

	"github.com/vvanpo/vind/types"
)

const prefixes = "kMGTPEZY"

type BytesFormat struct {
	// TODO Locale for , or .

	// Print the number of bytes without prefix multipliers like kB or GB
	NoPrefix bool
	// With prefixes enabled, toggles whether the prefix multiples should be
	// 1000 or 1024 (e.g. kilobyte "kB" or kibibyte "KiB")
	Base10 bool
}

func (f BytesFormat) Render(b types.Bytes) string {
	if f.NoPrefix {
		return fmt.Sprintf("%d B", b)
	}

	var n int

	if f.Base10 {
		n = int(math.Log10(float64(b))) / 3
	} else {
		n = int(math.Log2(float64(b))) / 10
	}

	if n == 0 {
		return fmt.Sprintf("%d B", b)
	}

	fr := float64(b)

	if f.Base10 {
		fr /= math.Pow10(n * 3)
	} else {
		fr = math.Ldexp(fr, -n*10)
	}

	return fmt.Sprintf("%.1f %sB", fr, f.prefix(n))
}

func (f BytesFormat) prefix(n int) string {
	i := "i"

	if f.Base10 {
		i = ""
	} else if n == 1 {
		return "Ki"
	}

	return string(prefixes[n-1]) + i
}
