package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.926 13.15a9 9 0 1 0-7.835 7.784"/><path d="M12 7v5l2 2m8 8l-5-5m0 5l5-5"/></g>`),
		g.Group(children),
	)
}