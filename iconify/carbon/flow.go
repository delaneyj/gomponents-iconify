package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 22.14V17a2 2 0 0 0-2-2h-8V9.86a4 4 0 1 0-2 0V15H7a2 2 0 0 0-2 2v5.14a4 4 0 1 0 2 0V17h18v5.14a4 4 0 1 0 2 0ZM8 26a2 2 0 1 1-2-2a2 2 0 0 1 2 2Zm6-20a2 2 0 1 1 2 2a2 2 0 0 1-2-2Zm12 22a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}