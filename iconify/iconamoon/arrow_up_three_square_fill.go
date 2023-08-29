package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpThreeSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V4Zm9.707 4.293a1 1 0 0 0-1.414 0l-3 3a1 1 0 1 0 1.414 1.414L11 11.414V15a1 1 0 1 0 2 0v-3.586l1.293 1.293a1 1 0 0 0 1.414-1.414l-3-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}