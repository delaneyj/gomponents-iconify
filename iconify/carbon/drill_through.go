package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrillThrough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 30A14.016 14.016 0 0 1 2 16h2A12 12 0 1 0 16 4V2a14 14 0 0 1 0 28Z"/><path fill="currentColor" d="M4 12v-2h4.586L2 3.414L3.414 2L10 8.586V4h2v8H4zm12-2v6h-6a6 6 0 1 0 6-6z"/>`),
		g.Group(children),
	)
}