package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuDotsBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M5 14a2 2 0 1 0-2-2"/><circle cx="12" cy="12" r="2"/><path stroke-linecap="round" d="M21 12a2 2 0 1 1-2-2"/></g>`),
		g.Group(children),
	)
}