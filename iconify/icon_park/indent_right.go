package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 9H6"/><path d="M29 19H6"/><path d="M29 29H6"/><path d="M37 19L42 24L37 29"/><path d="M42 39H6"/></g>`),
		g.Group(children),
	)
}