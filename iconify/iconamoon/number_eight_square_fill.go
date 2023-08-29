package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberEightSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-14ZM12 8a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm3 1c0 .675-.223 1.299-.6 1.8a4 4 0 1 1-4.8 0A3 3 0 1 1 15 9Zm-5 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}