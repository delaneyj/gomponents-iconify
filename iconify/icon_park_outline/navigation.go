package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Navigation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M24.5 4L9 44l15.5-9.09L40 44L24.5 4Z"/>`),
		g.Group(children),
	)
}