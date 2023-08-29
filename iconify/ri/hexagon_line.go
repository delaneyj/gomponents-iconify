package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 2.5L23 12l-5.5 9.5h-11L1 12l5.5-9.5h11Zm-1.153 2H7.653L3.311 12l4.342 7.5h8.694l4.342-7.5l-4.342-7.5Z"/>`),
		g.Group(children),
	)
}