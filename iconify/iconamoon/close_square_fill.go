package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V4Zm6.293 5.293a1 1 0 0 1 1.414 0L12 10.586l1.293-1.293a1 1 0 1 1 1.414 1.414L13.414 12l1.293 1.293a1 1 0 0 1-1.414 1.414L12 13.414l-1.293 1.293a1 1 0 0 1-1.414-1.414L10.586 12l-1.293-1.293a1 1 0 0 1 0-1.414Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}