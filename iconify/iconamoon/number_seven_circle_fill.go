package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSevenCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm7-6a1 1 0 0 0 0 2h4.523l-3.451 8.629a1 1 0 0 0 1.857.742l4-10A1 1 0 0 0 15 6H9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}