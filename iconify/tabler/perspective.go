package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Perspective(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6.141 4.163l12 1.714a1 1 0 0 1 .859.99v10.266a1 1 0 0 1-.859.99l-12 1.714A1 1 0 0 1 5 18.847V5.153a1 1 0 0 1 1.141-.99z"/>`),
		g.Group(children),
	)
}