package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClickToFold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M27 9V21H39"/><path d="M21 39V27H9"/><path d="M27 21L42 6"/><path d="M21 27L6 42"/></g>`),
		g.Group(children),
	)
}