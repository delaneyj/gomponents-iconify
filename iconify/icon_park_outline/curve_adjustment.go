package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurveAdjustment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M44 4H4v40h40V4Z"/><path stroke-linecap="round" d="M38 10c-6 0-11 4-14 14s-8 14-14 14"/></g>`),
		g.Group(children),
	)
}