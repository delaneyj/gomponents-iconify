package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalSpacingBetweenItems(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M8 6V12H40V6"/><path d="M14 24H34"/><path stroke-linejoin="round" d="M8 42V36H40V42"/></g>`),
		g.Group(children),
	)
}