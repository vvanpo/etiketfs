package local

import (
	"github.com/vvanpo/vind"
	"github.com/vvanpo/vind/client/driver"
)

type Driver struct {
	fs *vind.Filesystem
}

func (d Driver) List(f vind.Filter, s vind.Sort, ids []driver.Identifier) (<-chan []any, error) {
	files, err := d.fs.Select(f, s)

	if err != nil {
		return nil, err
	}

	out := make(chan []any)

	go func() {
		for f := range files {
			props := make([]any, len(ids))

			for i, id := range ids {
				if id.Parameter != nil {
					props[i], _ = f.Property(id.Group, id.Name, id.Parameter)
				} else {
					props[i], _ = f.Property(id.Group, id.Name)
				}
			}
		}

		close(out)
	}()

	return out, nil
}
