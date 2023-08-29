package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 0h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2Zm0 4v8h10V4H2Zm5 5.414l-1.414 1.414a1 1 0 1 1-1.414-1.414L5.586 8L4.172 6.586a1 1 0 1 1 1.414-1.414L7 6.586l1.414-1.414a1 1 0 1 1 1.414 1.414L8.414 8l1.414 1.414a1 1 0 0 1-1.414 1.414L7 9.414Z"/>`),
		g.Group(children),
	)
}