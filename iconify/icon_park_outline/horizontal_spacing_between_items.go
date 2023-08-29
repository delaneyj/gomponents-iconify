package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalSpacingBetweenItems(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 40h6V8H6"/><path d="M24 34V14"/><path stroke-linejoin="round" d="M42 40h-6V8h6"/></g>`),
		g.Group(children),
	)
}