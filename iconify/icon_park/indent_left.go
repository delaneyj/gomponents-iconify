package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 9H42"/><path d="M19 19H42"/><path d="M19 29H42"/><path d="M11 19L6 24L11 29"/><path d="M6 39H42"/></g>`),
		g.Group(children),
	)
}