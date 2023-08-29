package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockCancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.997 12.25a9 9 0 1 0-8.718 8.745M16 19a3 3 0 1 0 6 0a3 3 0 1 0-6 0m1 2l4-4"/><path d="M12 7v5l2 2"/></g>`),
		g.Group(children),
	)
}