package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tailoring(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M12 4V36H44"/><path stroke-linejoin="round" d="M20 12H36V28"/><path d="M12 12H4"/><path d="M36 44V36"/></g>`),
		g.Group(children),
	)
}