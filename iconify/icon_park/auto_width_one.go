package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoWidthOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M6 7H42"/><path d="M8 24H40"/><path stroke-linejoin="round" d="M13.9907 30L8 24.0046L14 18"/><path stroke-linejoin="round" d="M34.0093 18L40 23.9954L34 30"/><path d="M6 41H42"/></g>`),
		g.Group(children),
	)
}