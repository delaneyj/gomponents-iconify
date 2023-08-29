package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wrench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.47 9.8A5 5 0 0 1 .2 3.22l3.95 3.95l2.82-2.83L3.03.41a5 5 0 0 1 6.4 6.68l10 10l-2.83 2.83L6.47 9.8z"/>`),
		g.Group(children),
	)
}