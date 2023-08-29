package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExitFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 3.001a1 1 0 1 1 0 2H6v13a1 1 0 0 0 1 1h8a1 1 0 1 1 0 2H7a3 3 0 0 1-3-3v-14a1 1 0 0 1 1-1h10Zm1.707 5.293A1 1 0 0 0 15 9v2H9a1 1 0 1 0 0 2h6v2a1 1 0 0 0 1.707.707l3-3a1 1 0 0 0 0-1.414l-3-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}