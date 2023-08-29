package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoClipOffTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 1 0-1.06 1.06l.702.702A3.24 3.24 0 0 0 2 6.25v11.5A3.25 3.25 0 0 0 5.25 21h13.5a3.24 3.24 0 0 0 1.024-.165l.945.945a.75.75 0 0 0 1.061-1.06L3.28 2.22Zm9.836 11.957l-2.634 1.45A1 1 0 0 1 9 14.75v-4.69l4.117 4.117Zm2.366-3.053c.563.31.667 1.027.311 1.487l6.071 6.072c.089-.296.136-.609.136-.933V6.25A3.25 3.25 0 0 0 18.75 3H6.182l6.688 6.688l2.612 1.436Z"/>`),
		g.Group(children),
	)
}