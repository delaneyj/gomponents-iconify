package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BombMinimalisticLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="9.5" cy="14.5" r="7.5"/><path stroke-linecap="round" d="m17 7l-2 2" opacity=".5"/><path stroke-linecap="round" d="m19.5 7.5l1 .5M16 3.5l.5 1M19 5l1-1"/></g>`),
		g.Group(children),
	)
}