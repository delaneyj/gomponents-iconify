package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaClosedF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 13.983l-6 1.582L6.234 2.207a1 1 0 1 1 1.533 0L14 15.565l-6-1.582v3.582a1 1 0 0 0 2 0a1 1 0 0 1 2 0a3 3 0 0 1-6 0v-3.582z"/>`),
		g.Group(children),
	)
}