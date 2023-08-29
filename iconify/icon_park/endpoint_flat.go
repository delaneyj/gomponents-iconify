package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EndpointFlat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M43 40H23.2857C20.0414 40 7 40 7 40L7 27M43 8H23.2857C14.2914 8 7 8 7 8L7 21M43 24H10"/><circle cx="7" cy="24" r="3"/></g>`),
		g.Group(children),
	)
}