package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m36 7l7 6.461L36 21"/><path d="M40 14H17.006C10.123 14 4.278 19.62 4.01 26.5C3.726 33.77 9.733 40 17.006 40h16.996"/></g>`),
		g.Group(children),
	)
}