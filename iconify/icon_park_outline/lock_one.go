package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="24" cy="30" r="14"/><path stroke-linejoin="round" d="M31 18v-7a7 7 0 1 0-14 0v7"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 26v8"/></g>`),
		g.Group(children),
	)
}