package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 8V40"/><path d="M25 8V40"/><path d="M6 24H25"/><path d="M34.2261 24L39.0001 19.0166V40"/></g>`),
		g.Group(children),
	)
}