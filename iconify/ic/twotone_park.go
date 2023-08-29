package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotonePark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.26 10h1.9l-3.15-4.5L8.88 10h1.81l-3.9 6h10.47z" opacity=".3"/><path fill="currentColor" d="M17 12h2L12 2L5.05 12H7l-3.9 6h6.92v4h3.95v-4H21l-4-6zM6.79 16l3.9-6H8.88l3.13-4.5l3.15 4.5h-1.9l4 6H6.79z"/>`),
		g.Group(children),
	)
}