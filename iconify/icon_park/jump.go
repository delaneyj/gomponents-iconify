package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jump(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M8 10L20 20.254V29.9683L10.8571 44"/><path stroke-linecap="round" stroke-linejoin="round" d="M40 10L28 20.254V29.9683L37.1429 44"/><circle cx="24" cy="8" r="4" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}