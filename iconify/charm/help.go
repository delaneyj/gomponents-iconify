package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Help(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M5.75 6.75c0-1 1-2 2.25-2s2.25 1.034 2.25 2c0 1.5-1.5 1.5-2.25 2m0 2.5v0"/></g>`),
		g.Group(children),
	)
}