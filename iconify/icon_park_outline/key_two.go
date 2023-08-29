package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="15" cy="33" r="8"/><path stroke-linecap="round" stroke-linejoin="round" d="m29 16l7 6m-16 4L36 8l7 6"/></g>`),
		g.Group(children),
	)
}