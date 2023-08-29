package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.943 13.016A9 9 0 1 0 12.028 21M16 22l5-5m0 4.5V17h-4.5"/><path d="M12 7v5l2 2"/></g>`),
		g.Group(children),
	)
}