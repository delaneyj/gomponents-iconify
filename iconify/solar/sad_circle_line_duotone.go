package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SadCircleLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="1.5" opacity=".5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M9 17c.85-.63 1.885-1 3-1s2.15.37 3 1"/><ellipse cx="15" cy="10.5" fill="currentColor" rx="1" ry="1.5"/><ellipse cx="9" cy="10.5" fill="currentColor" rx="1" ry="1.5"/></g>`),
		g.Group(children),
	)
}