package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DividingLineOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 24H44"/><path d="M8 10H12"/><path d="M20 10H28"/><path d="M36 10H40"/><path d="M8 38H12"/><path d="M20 38H28"/><path d="M36 38H40"/></g>`),
		g.Group(children),
	)
}