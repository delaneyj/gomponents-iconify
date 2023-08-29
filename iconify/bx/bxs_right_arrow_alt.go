package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsRightArrowAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M19 12l-7-6v5H6v2h6v5z" fill="currentColor"/>`),
		g.Group(children),
	)
}