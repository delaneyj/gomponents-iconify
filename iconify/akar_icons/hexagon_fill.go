package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.036 1.53a2 2 0 0 1 1.928 0l8 4.4A2 2 0 0 1 22 7.683v8.634a2 2 0 0 1-1.036 1.753l-8 4.4a2 2 0 0 1-1.928 0l-8-4.4A2 2 0 0 1 2 16.317V7.683A2 2 0 0 1 3.036 5.93l8-4.4Z"/>`),
		g.Group(children),
	)
}