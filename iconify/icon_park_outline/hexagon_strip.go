package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonStrip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 4h10v11.34l9.82-5.67l5 8.66L34 24l9.82 5.67l-5 8.66L29 32.66V44H19V32.66l-9.82 5.67l-5-8.66L14 24l-9.82-5.67l5-8.66L19 15.34V4Z"/>`),
		g.Group(children),
	)
}