package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="13.715" cy="13.714" r="6.857" fill="#2F88FF"/><circle cx="34.286" cy="34.285" r="6.857" fill="#2F88FF"/><path d="M24.001 44C12.9553 44 4.00098 35.0457 4.00098 24L10.6676 27.3333"/><path d="M24.001 4C35.0467 4 44.001 12.9543 44.001 24L37.3343 20.6667"/></g>`),
		g.Group(children),
	)
}