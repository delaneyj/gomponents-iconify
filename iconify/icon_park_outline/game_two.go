package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m20 15l4 4l4-4V4h-8v11Zm0 18l4-4l4 4v11h-8V33Zm13-5l-4-4l4-4h11v8H33Zm-18-8l4 4l-4 4H4v-8h11Z"/>`),
		g.Group(children),
	)
}