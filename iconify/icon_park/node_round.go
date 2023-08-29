package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NodeRound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 42V27M42 24H27"/><circle cx="24" cy="24" r="3"/><path d="M42 6H24C14.0589 6 6 14.0589 6 24V42"/></g>`),
		g.Group(children),
	)
}