package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chisel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 14l1.5 1.5m2.847.075l2.08 2.079a1.96 1.96 0 0 1-2.773 2.772l-2.08-2.079a1.96 1.96 0 0 1 2.773-2.772zM3 6l3-3l7.414 7.414A2 2 0 0 1 14 11.828V14h-2.172a2 2 0 0 1-1.414-.586L3 6z"/>`),
		g.Group(children),
	)
}