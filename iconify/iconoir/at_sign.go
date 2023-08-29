package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.5 19.125A9 9 0 1 1 21 12c0 5.5-6 5.5-6 2V8"/><path d="M15 12v-1.5C15 9.12 13.657 8 12 8s-3 1.12-3 2.5V12m6 0v1.5c0 1.38-1.343 2.5-3 2.5s-3-1.12-3-2.5V12"/></g>`),
		g.Group(children),
	)
}