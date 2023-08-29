package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCrossBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="10" cy="6" r="4"/><path stroke-linecap="round" d="M20.414 11.414L19 10m0 0l-1.414-1.414M19 10l1.414-1.414M19 10l-1.414 1.414M17.997 18c.003-.164.003-.331.003-.5c0-2.485-3.582-4.5-8-4.5s-8 2.015-8 4.5S2 22 10 22c2.231 0 3.84-.157 5-.437"/></g>`),
		g.Group(children),
	)
}