package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsMapAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M9 6.882l-7-3.5v13.236l7 3.5l6-3l7 3.5V7.382l-7-3.5l-6 3zM15 15l-6 3V9l6-3v9z" fill="currentColor"/>`),
		g.Group(children),
	)
}