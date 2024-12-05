package main

import (
	"fmt"
	"time"

	"github.com/vvanpo/vind/client/values"
	"github.com/vvanpo/vind/types"
)

type formatters struct {
	bytes     values.BytesFormat
	stringset values.StringSetFormat
	time      values.TimeFormat
}

func defaults() formatters {
	return formatters{
		time: values.TimeFormat{Location: time.Local},
	}
}

func (fs *formatters) format(val any) string {
	switch v := val.(type) {
	case types.Bytes:
		return fs.bytes.Format(v)
	case types.StringSet:
		return fs.stringset.Format(v)
	case types.Time:
		return fs.time.Format(v)
	case nil:
		return ""
	}

	return fmt.Sprint(val)
}
