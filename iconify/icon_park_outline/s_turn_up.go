package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func STurnUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M24 31v4c0 4-3 7-7 7s-7-3-7-7V16m28 26V13c0-4-3-7-7-7s-7 3-7 7v18"/><path stroke-linecap="round" stroke-linejoin="round" d="m33 37l5 5l5-5"/><circle cx="10" cy="11" r="5" transform="rotate(-180 10 11)"/></g>`),
		g.Group(children),
	)
}