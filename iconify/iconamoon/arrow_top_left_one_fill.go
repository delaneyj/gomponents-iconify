package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopLeftOneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 6H7a1 1 0 0 0-1 1v10a1 1 0 0 0 1.707.707L12 13.414l4.293 4.293a1 1 0 0 0 1.414-1.414L13.414 12l4.293-4.293A1 1 0 0 0 17 6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}