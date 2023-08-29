package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuFold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 11H40"/><path d="M8 24H40"/><path d="M8 37H40"/><path d="M13.6568 29.6568L8 23.9999L13.6568 18.343"/></g>`),
		g.Group(children),
	)
}