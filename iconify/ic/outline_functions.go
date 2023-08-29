package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineFunctions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 4H6v2l6.5 6L6 18v2h12v-3h-7l5-5l-5-5h7V4z"/>`),
		g.Group(children),
	)
}