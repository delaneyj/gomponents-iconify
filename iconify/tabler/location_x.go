package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13.5 21l-.224-.448L10 14l-7-3.5a.55.55 0 0 1 0-1L21 3l-3.622 10.03M22 22l-5-5m0 5l5-5"/>`),
		g.Group(children),
	)
}