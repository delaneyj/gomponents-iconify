package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarFourPoints(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1L9 9l-8 3l8 3l3 8l3-8l8-3l-8-3l-3-8Z"/>`),
		g.Group(children),
	)
}