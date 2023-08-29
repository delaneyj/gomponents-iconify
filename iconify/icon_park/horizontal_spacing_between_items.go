package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalSpacingBetweenItems(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 40H12L12 8H6"/><path d="M24 34V14"/><path stroke-linejoin="round" d="M42 40H36V8H42"/></g>`),
		g.Group(children),
	)
}