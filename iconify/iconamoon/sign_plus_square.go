package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignPlusSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v6m-3-3h6M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4Z"/>`),
		g.Group(children),
	)
}