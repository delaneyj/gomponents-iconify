package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CactusFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 216a8 8 0 0 1-8 8H40a8 8 0 0 1 0-16h48v-64h-4a68.07 68.07 0 0 1-68-68a28 28 0 0 1 56 0a12 12 0 0 0 12 12h4V56a40 40 0 0 1 80 0v72h4a12 12 0 0 0 12-12a28 28 0 0 1 56 0a68.07 68.07 0 0 1-68 68h-4v24h48a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}