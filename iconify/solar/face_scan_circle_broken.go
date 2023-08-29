package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceScanCircleBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M9 16c.85.63 1.885 1 3 1s2.15-.37 3-1"/><ellipse cx="15" cy="10.5" fill="currentColor" rx="1" ry="1.5"/><ellipse cx="9" cy="10.5" fill="currentColor" rx="1" ry="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M2.2 10A10.012 10.012 0 0 1 10 2.2M2.2 14a10.012 10.012 0 0 0 7.8 7.8M21.8 10A10.012 10.012 0 0 0 14 2.2M21.8 14a10.012 10.012 0 0 1-7.8 7.8"/></g>`),
		g.Group(children),
	)
}