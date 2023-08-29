package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpOneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19 13a1 1 0 0 0 .707-1.707l-7-7a1 1 0 0 0-1.414 0l-7 7A1 1 0 0 0 5 13h6v6a1 1 0 1 0 2 0v-6h6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}