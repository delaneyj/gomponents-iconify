package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Document(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 18h12V6h-4V2H4v16zm-2 1V0h12l4 4v16H2v-1z"/>`),
		g.Group(children),
	)
}