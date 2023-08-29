package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m17.704 6.58l-6.972 8.086a.967.967 0 0 1-1.463 0L2.296 6.58C1.76 5.96 1.967 5 2.791 5h14.42c.821 0 1.029.96.493 1.58Z"/>`),
		g.Group(children),
	)
}