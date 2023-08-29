package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberEightCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><circle cx="12" cy="14" r="3"/><circle cx="12" cy="9" r="2"/></g>`),
		g.Group(children),
	)
}