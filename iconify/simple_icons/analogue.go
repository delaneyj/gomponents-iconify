package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analogue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.468 12.804a5.145 5.145 0 1 0-.644 10.27a5.145 5.145 0 0 0 .644-10.27zm17.841 2.562L16.45 3.484a5.146 5.146 0 0 0-8.912 5.15l6.86 11.878a5.148 5.148 0 0 0 7.031 1.885a5.146 5.146 0 0 0 1.881-7.031z"/>`),
		g.Group(children),
	)
}