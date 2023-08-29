package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsHorizontalCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 0 1 10 10a10 10 0 0 1-10 10A10 10 0 0 1 2 12A10 10 0 0 1 12 2m0 2a8 8 0 0 0-8 8a8 8 0 0 0 8 8a8 8 0 0 0 8-8a8 8 0 0 0-8-8m0 6.5a1.5 1.5 0 0 1 1.5 1.5a1.5 1.5 0 0 1-1.5 1.5a1.5 1.5 0 0 1-1.5-1.5a1.5 1.5 0 0 1 1.5-1.5m-4.5 0A1.5 1.5 0 0 1 9 12a1.5 1.5 0 0 1-1.5 1.5A1.5 1.5 0 0 1 6 12a1.5 1.5 0 0 1 1.5-1.5m9 0A1.5 1.5 0 0 1 18 12a1.5 1.5 0 0 1-1.5 1.5A1.5 1.5 0 0 1 15 12a1.5 1.5 0 0 1 1.5-1.5Z"/>`),
		g.Group(children),
	)
}