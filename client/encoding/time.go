package encoding

import (
	"time"

	"github.com/vvanpo/vind/types"
)

type TimeFormat struct {
	*time.Location
	Absolute bool // TODO
}

func (f TimeFormat) Render(t types.Time) string {
	loc := time.UTC

	if f.Location != nil {
		loc = f.Location
	}

	p := time.Unix(int64(t), 0).In(loc)

	return p.Format(time.DateTime)
}
