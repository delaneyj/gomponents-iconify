package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsVerticalCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 12A10 10 0 0 1 12 2a10 10 0 0 1 10 10a10 10 0 0 1-10 10A10 10 0 0 1 2 12m8.5 0a1.5 1.5 0 0 0 1.5 1.5a1.5 1.5 0 0 0 1.5-1.5a1.5 1.5 0 0 0-1.5-1.5a1.5 1.5 0 0 0-1.5 1.5m0 5.5A1.5 1.5 0 0 0 12 19a1.5 1.5 0 0 0 1.5-1.5A1.5 1.5 0 0 0 12 16a1.5 1.5 0 0 0-1.5 1.5m0-11A1.5 1.5 0 0 0 12 8a1.5 1.5 0 0 0 1.5-1.5A1.5 1.5 0 0 0 12 5a1.5 1.5 0 0 0-1.5 1.5Z"/>`),
		g.Group(children),
	)
}