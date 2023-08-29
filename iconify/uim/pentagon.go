package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pentagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.562 21.56H6.437a1 1 0 0 1-.95-.692l-3.438-10.58a.999.999 0 0 1 .363-1.117l9-6.54a.996.996 0 0 1 1.176 0l9 6.54a.999.999 0 0 1 .363 1.117l-3.437 10.58a1 1 0 0 1-.952.692Z"/>`),
		g.Group(children),
	)
}