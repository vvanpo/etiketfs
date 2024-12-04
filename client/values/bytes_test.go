package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vvanpo/vind/types"
)

func TestNoPrefix(t *testing.T) {
	out := BytesFormat{NoPrefix: true}.Format(types.Bytes(14063514))

	assert.Equal(t, "14063514 B", out)
}

func TestBase10Small(t *testing.T) {
	out := BytesFormat{}.Format(types.Bytes(963))

	assert.Equal(t, "963 B", out)
}

func TestBase10Kilo(t *testing.T) {
	out := BytesFormat{Base10: true}.Format(types.Bytes(490502))

	assert.Equal(t, "490.5 kB", out)
}

func TestBase10Giga(t *testing.T) {
	out := BytesFormat{Base10: true}.Format(types.Bytes(513602659821))

	assert.Equal(t, "513.6 GB", out)
}

func TestBase2Small(t *testing.T) {
	out := BytesFormat{}.Format(types.Bytes(1017))

	assert.Equal(t, "1017 B", out)
}

func TestBase2Kibi(t *testing.T) {
	out := BytesFormat{}.Format(types.Bytes(490502))

	assert.Equal(t, "479.0 KiB", out)
}

func TestBase2Gigi(t *testing.T) {
	out := BytesFormat{}.Format(types.Bytes(513602659821))

	assert.Equal(t, "478.3 GiB", out)
}
