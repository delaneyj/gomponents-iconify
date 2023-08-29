package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="24" cy="30" r="14"/><path stroke-linecap="round" stroke-linejoin="round" d="M31 12v-1a7 7 0 0 0-7-7v0a7 7 0 0 0-7 7v6m7 9v8"/></g>`),
		g.Group(children),
	)
}