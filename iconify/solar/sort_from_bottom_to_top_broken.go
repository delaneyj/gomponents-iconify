package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortFromBottomToTopBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 8h9m-7 5h7m-5 5h5"/><path stroke-linejoin="round" d="M17 20V4l3 4"/></g>`),
		g.Group(children),
	)
}