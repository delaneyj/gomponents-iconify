package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottlesTwoDash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 9.051V6a1 1 0 0 0-1-1h-3v2h2v3.02s2 1.124 2 3.48V25h-4v2h5a1 1 0 0 0 1-1V13.5a5.93 5.93 0 0 0-2-4.449zM16 25h-3v2h4a1 1 0 0 0 1-1v-4h-2v3zm-8-3H6v4a1 1 0 0 0 1 1h4v-2H8v-3zm8-6h2v4h-2zM6 16h2v4H6zm10-6.949V6a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v3.051A5.93 5.93 0 0 0 6 13.5v.5h2v-.5c0-2.356 2-3.48 2-3.48V7h4v3.02s2 1.124 2 3.48v.5h2v-.5a5.93 5.93 0 0 0-2-4.449z"/>`),
		g.Group(children),
	)
}