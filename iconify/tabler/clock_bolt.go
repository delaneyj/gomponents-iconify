package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockBolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.984 12.53a9 9 0 1 0-7.552 8.355"/><path d="M12 7v5l3 3m4 1l-2 3h4l-2 3"/></g>`),
		g.Group(children),
	)
}