package coordinate_supplier

import (
	"fmt"
	"sync/atomic"
)

type coordinateSupplierAtomic struct {
	coordinates []Coordinate
	at          uint64
	done        uint64
	repeat      bool
	order       Order
}

// NewCoordinateSupplierAtomic returns a CoordinateSupplier synchronized with atomic.AddUint64.
// It is the fastest implementation but some coordinates could be received slightly out-of-order when called concurrently.
func NewCoordinateSupplierAtomic(opts CoordinateSupplierOptions) (CoordinateSupplier, error) {
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

	cs := &coordinateSupplierAtomic{
		repeat:      opts.Repeat,
		coordinates: coords,
	}

	return cs, nil
}

// Next returns the next coordinate to be supplied.
// It may be possible to receive some coordinates slightly out of order when called concurrently.
func (c *coordinateSupplierAtomic) Next() (x, y int, done bool) {
	// check if already done
	if atomic.LoadUint64(&c.done) > 0 {
		// already done
		return 0, 0, true
	}

	// concurrent-safe and in-order get the next element index
	atNow := atomic.AddUint64(&c.at, 1) - 1

	// check if now done
	if !c.repeat && atNow >= uint64(len(c.coordinates)) {
		// mark as done
		atomic.AddUint64(&c.done, 1)
		return 0, 0, true
	}

	// if repeating past the end, clamp to the current remainder position
	atNowClamped := atNow % uint64(len(c.coordinates))

	// return matching coordinate (by now the timing may be slightly out-of-order)
	return c.coordinates[atNowClamped].X, c.coordinates[atNowClamped].Y, false
}
