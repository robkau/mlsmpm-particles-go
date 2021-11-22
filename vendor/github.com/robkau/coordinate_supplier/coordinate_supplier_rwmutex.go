package coordinate_supplier

import (
	"fmt"
	"sync"
)

type coordinateSupplierRWMutex struct {
	coordinates []Coordinate
	at          int
	repeat      bool
	order       Order
	rw          sync.RWMutex
}

// NewCoordinateSupplierRWMutex returns a CoordinateSupplier synchronized with sync.RWMutex.
// It blocks more and is slower than NewCoordinateSupplierAtomic, but the coordinates are guaranteed to be handed out strictly in-order when used concurrently.
func NewCoordinateSupplierRWMutex(opts CoordinateSupplierOptions) (CoordinateSupplier, error) {
	if opts.Width < 1 {
		return nil, fmt.Errorf("minimum width is 1")
	}
	if opts.Height < 1 {
		return nil, fmt.Errorf("minimum height is 1")
	}
	coords, err := MakeCoordinateList(opts.Width, opts.Height, opts.Order)
	if err != nil {
		return nil, fmt.Errorf("failed make coordinate list: %w", err)
	}

	cs := &coordinateSupplierRWMutex{
		repeat:      opts.Repeat,
		rw:          sync.RWMutex{},
		coordinates: coords,
	}
	return cs, nil
}

// Next returns the next coordinate to be supplied.
func (c *coordinateSupplierRWMutex) Next() (x, y int, done bool) {
	c.rw.Lock()
	defer c.rw.Unlock()

	if c.at >= len(c.coordinates) {
		if c.repeat {
			c.at = 0
		} else {
			return 0, 0, true
		}
	}

	defer func() { c.at++ }()
	return c.coordinates[c.at].X, c.coordinates[c.at].Y, false
}
