package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Editing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="13" cy="35" r="7"/><circle cx="35" cy="35" r="7"/><path stroke-linecap="round" stroke-linejoin="round" d="m8 6l24 22m8-22L16 28"/></g>`),
		g.Group(children),
	)
}