package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRoundRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 29H6v14h14V29Zm4-25l10 17H14L24 4Zm12 40a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/>`),
		g.Group(children),
	)
}