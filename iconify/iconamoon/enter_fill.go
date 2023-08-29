package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 3a1 1 0 0 0 0 2h9v13a1 1 0 0 1-1 1H9a1 1 0 0 0 0 2h8a3 3 0 0 0 3-3V4a1 1 0 0 0-1-1H9Zm3.707 5.293A1 1 0 0 0 11 9v2H5a1 1 0 0 0 0 2h6v2a1 1 0 0 0 1.707.707l3-3a1 1 0 0 0 0-1.414l-3-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}