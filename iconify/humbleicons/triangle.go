package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Triangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M11.125 2.584a1 1 0 0 1 1.75 0l8.805 15.932A1 1 0 0 1 20.805 20H3.195a1 1 0 0 1-.875-1.484l8.805-15.932z"/>`),
		g.Group(children),
	)
}