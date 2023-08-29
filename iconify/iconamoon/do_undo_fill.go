package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoUndoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.383 5.076A1 1 0 0 1 9 6v2h7a5 5 0 0 1 0 10H5a1 1 0 1 1 0-2h11a3 3 0 1 0 0-6H9v2a1 1 0 0 1-1.707.707l-3-3a1 1 0 0 1 0-1.414l3-3a1 1 0 0 1 1.09-.217Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}