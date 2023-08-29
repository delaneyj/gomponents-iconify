package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircuitComposer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M18 9v6h-4V9h-2v14h2v-6h4v6h2V9h-2z" fill="currentColor"/><path d="M30 15h-4V6a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v9H2v2h4v9a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-9h4zM8 26V6h16v20z" fill="currentColor"/>`),
		g.Group(children),
	)
}