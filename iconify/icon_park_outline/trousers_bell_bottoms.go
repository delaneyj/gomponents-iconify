package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrousersBellBottoms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m16 12l2-8h12l2 8v12l4 17l-12 3l-12-3l4-17V12Zm8 32V16"/><path d="m12 41l12 3l12-3"/></g>`),
		g.Group(children),
	)
}