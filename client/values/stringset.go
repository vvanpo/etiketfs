package values

import "github.com/vvanpo/vind/types"

type StringSetFormat struct{}

func (ss StringSetFormat) Format(t types.StringSet) string {
	if len(t) == 0 {
		return ""
	}

	var out string

	for k := range t {
		out += ", " + k
	}

	return out[2:]
}
