package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 9h18M24 19h18M6 29h36M6 39h36M6 9.766a1 1 0 0 1 1.514-.857l7.057 4.233a1 1 0 0 1 0 1.716L7.515 19.09A1 1 0 0 1 6 18.234V9.766Z"/>`),
		g.Group(children),
	)
}