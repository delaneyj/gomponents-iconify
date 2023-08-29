package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomLeftThreeSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V4Zm6 6v4a1 1 0 0 0 1 1h4a1 1 0 1 0 0-2h-1.586l2.293-2.293a1 1 0 0 0-1.414-1.414L11 11.586V10a1 1 0 1 0-2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}