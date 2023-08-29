package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16h6v2H9v-2zm13 1h-3v3a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-3H2V9a2 2 0 0 1 2-2h1V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v2h1a2 2 0 0 1 2 2v8zM7 7h10V5H7v2zm10 7H7v6h10v-6zm3-3.5a1.5 1.5 0 1 0-3.001.001A1.5 1.5 0 0 0 20 10.5z"/>`),
		g.Group(children),
	)
}