package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 2H6L2 8l10 14L22 8l-4-6M4.43 8l2.64-4h9.86l2.64 4L12 18.56L4.43 8Z"/>`),
		g.Group(children),
	)
}