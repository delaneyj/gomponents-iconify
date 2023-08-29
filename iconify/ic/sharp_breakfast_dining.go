package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpBreakfastDining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.85 3H6.14C4.15 3 2.36 4.39 2.05 6.36c-.27 1.75.59 3.29 1.95 4.09V21h16V10.45a3.996 3.996 0 0 0 1.95-4.11C21.63 4.38 19.83 3 17.85 3zm-1.44 10L12 17.42L7.59 13L12 8.59L16.41 13z"/>`),
		g.Group(children),
	)
}