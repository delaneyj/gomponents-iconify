package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkPing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18v-2h6.5L2 7.5l1.4-1.4l8.6 8.575l5.2-5.2q-.1-.225-.15-.462T17 8.5q0-1.05.725-1.775T19.5 6q1.05 0 1.775.725T22 8.5q0 1.05-.725 1.775T19.5 11q-.225 0-.438-.037t-.412-.113L13.5 16H20v2H4Z"/>`),
		g.Group(children),
	)
}