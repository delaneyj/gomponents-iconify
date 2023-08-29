package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.931 13.111a9 9 0 1 0-9.453 7.874M20 21l2-2l-2-2m-3 0l-2 2l2 2"/><path d="M12 7v5l2 2"/></g>`),
		g.Group(children),
	)
}