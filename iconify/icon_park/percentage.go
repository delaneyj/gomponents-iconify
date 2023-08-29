package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Percentage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="11" cy="11" r="5" fill="#2F88FF"/><circle cx="37" cy="37" r="5" fill="#2F88FF"/><path d="M42 6L6 42"/></g>`),
		g.Group(children),
	)
}