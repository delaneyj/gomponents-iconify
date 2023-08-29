package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dvr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 15v-4a2 2 0 0 0-2-2h-6v14h2v-6h1.48l2.34 6H30l-2.33-6H28a2 2 0 0 0 2-2zm-6-4h4v4h-4zm-6-2l-2 13l-2-13h-2l2.52 14h2.96L20 9h-2zM6 23H2V9h4a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4zm-2-2h2a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2H4z"/>`),
		g.Group(children),
	)
}