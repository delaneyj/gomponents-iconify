package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaComponent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m17 12l7-7l7 7l-7 7l-7-7Zm0 24l7-7l7 7l-7 7l-7-7Zm12-12l7-7l7 7l-7 7l-7-7ZM5 24l7-7l7 7l-7 7l-7-7Z"/>`),
		g.Group(children),
	)
}