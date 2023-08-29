package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSevenSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-14ZM9 6a1 1 0 0 0 0 2h4.523l-3.451 8.629a1 1 0 0 0 1.857.742l4-10A1 1 0 0 0 15 6H9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}