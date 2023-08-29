package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 9.5A3.5 3.5 0 0 0 13 6H8.5a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1H13a3.49 3.49 0 0 0 2.44-6a3.5 3.5 0 0 0 1.06-2.5ZM13 16H9.5v-3H13a1.5 1.5 0 0 1 0 3Zm0-5H9.5V8H13a1.5 1.5 0 0 1 0 3Z"/>`),
		g.Group(children),
	)
}