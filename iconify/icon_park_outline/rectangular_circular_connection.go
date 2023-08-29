package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangularCircularConnection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M12 19a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm5 12H7v10h10V31Z"/><path stroke-linecap="round" d="M25.68 13H42v23H25"/></g>`),
		g.Group(children),
	)
}