package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderOutside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 11H7v2h2m4 2h-2v2h2m6 2H5V5h14M3 21h18V3H3m14 8h-2v2h2m-4-2h-2v2h2m0-6h-2v2h2V7Z"/>`),
		g.Group(children),
	)
}