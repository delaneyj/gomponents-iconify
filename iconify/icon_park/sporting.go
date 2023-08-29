package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sporting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="24" cy="8" r="4" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M7 18H19V34"/><path stroke-linecap="round" stroke-linejoin="round" d="M41 18H29V44"/></g>`),
		g.Group(children),
	)
}