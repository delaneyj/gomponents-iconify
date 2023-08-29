package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFourSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-14Zm8.461 3.274a1 1 0 1 0-1.923-.55l-1.271 4.45A3 3 0 0 0 11.152 15H13.5v2a1 1 0 1 0 2 0v-7a1 1 0 1 0-2 0v3h-2.348a1 1 0 0 1-.962-1.275l1.271-4.45Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}