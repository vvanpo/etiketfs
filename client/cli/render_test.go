package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vvanpo/vind"
	"github.com/vvanpo/vind/client/driver"
	"github.com/vvanpo/vind/types"
)

type dummy struct{}

func (d *dummy) List(filter vind.Filter, output []driver.Identifier) <-chan []any {
	c := make(chan []any)

	go func() {
		c <- []any{
			types.StringSet(map[string]struct{}{"jpeg": {}}),
			types.Time(time.Date(2024, 12, 6, 14, 02, 25, 0, time.Local).Unix()),
			types.Bytes(490502),
			nil,
		}
		c <- []any{
			types.StringSet(map[string]struct{}{"utf8": {}, "ascii": {}}),
			types.Time(time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local).Unix()),
			types.Bytes(230),
			230,
		}
		close(c)
	}()

	return c
}

func TestHelloWorld(t *testing.T) {
	fmts := defaults()
	d := new(dummy)
	files := d.List(vind.Filter{}, nil)

	props := render(fmts, files)

	assert.Equal(t, []string{"jpeg", "2024-12-06 14:02:25", "479.0 KiB", ""}, <-props)
	assert.Equal(t, []string{"utf8, ascii", "2006-01-02 15:04:05", "230 B", "230"}, <-props)
	_, ok := <-props
	assert.False(t, ok)
}
