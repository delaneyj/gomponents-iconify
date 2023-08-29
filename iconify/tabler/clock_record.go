package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockRecord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12.3a9 9 0 1 0-8.683 8.694"/><path d="M12 7v5l2 2m2 5a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/></g>`),
		g.Group(children),
	)
}