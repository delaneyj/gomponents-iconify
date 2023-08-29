package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatOnce(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 17H8c-1.667 0-5-1-5-5m5-5h8c1.667 0 5 1 5 5c0 1.494-.465 2.57-1.135 3.331"/><path d="M14.5 14.5L17 17l-2.5 2.5M4 8V3L2 4"/></g>`),
		g.Group(children),
	)
}