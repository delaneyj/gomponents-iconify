package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignRadicalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.026 4.773A1 1 0 0 1 8 4h13a1 1 0 1 1 0 2H8.794l-3.32 14.227a1 1 0 0 1-1.932.06l-1.5-5a1 1 0 1 1 1.916-.574l.421 1.404L7.026 4.773Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}