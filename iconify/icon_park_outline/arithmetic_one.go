package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArithmeticOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M28 31.5h14m-14 8h14M7.343 40.657l11.314-11.314m-11.314 0l11.314 11.314M28 12.5h14m-36 0h14m-7-7v14"/>`),
		g.Group(children),
	)
}