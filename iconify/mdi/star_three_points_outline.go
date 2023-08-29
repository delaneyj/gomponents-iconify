package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarThreePointsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 9.5l1.2 4l2.8 3l-4-.9l-4.1.9l2.8-3l1.3-4m0-6.9l-3 9.8l-7 7.5l10-2.3L22 20l-7-7.5l-3-9.9Z"/>`),
		g.Group(children),
	)
}