package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuUnfold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 11H40"/><path d="M8 24H42"/><path d="M8 37H40"/><path d="M36.3433 29.6568L42.0001 23.9999L36.3433 18.343"/></g>`),
		g.Group(children),
	)
}