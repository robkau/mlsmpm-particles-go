package coordinate_supplier

/* Order determines how coordinates should be handed out:
 - in ascending order: 1, 2, 3, ...
 - in descending order: 3, 2, 1, ...
 - in random order: 2, 1, 3, ...
Coordinate system origin can be imagined in the bottom left. When ascending order, elements will be handed out left-to-right and then bottom-to-top.
These are the first 9 points handed out when in ascending order for a 3x3 grid:

   y2 || (7) 0,2     (8) 1,2     (9) 2,2
      ||
   y1 || (4) 0,1     (5) 1,1     (6) 2,1
      ||
   y0 || (1) 0,0     (2) 1,0     (3) 2,0
	   =================================
             x0          x1          x2


*/
type Order uint

const (
	Asc Order = iota
	Desc
	Random
)

func OrderToString(o Order) string {
	switch o {
	case Asc:
		return "Asc"
	case Desc:
		return "Desc"
	case Random:
		return "Random"
	default:
		return ""
	}
}
