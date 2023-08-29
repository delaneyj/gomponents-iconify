package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShiJumping(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 3a1 1 0 1 0 2 0a1 1 0 0 0-2 0m6 14.5L12 13V7l5 4M7 17.5l5-4.5"/><path d="m15.103 21.58l6.762-14.502a2 2 0 0 0-.968-2.657m-12 17.159L2.135 7.077a2 2 0 0 1 .968-2.657M7 11l5-4"/></g>`),
		g.Group(children),
	)
}