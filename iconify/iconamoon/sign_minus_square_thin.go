package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignMinusSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 12h6M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4Z"/>`),
		g.Group(children),
	)
}