package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinPeopleNearby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="15" cy="10" r="4"/><circle cx="33" cy="10" r="4"/><path d="M10 20H20L18 42H12L10 20Z"/><path d="M28 20H38L36 42H30L28 20Z"/></g>`),
		g.Group(children),
	)
}