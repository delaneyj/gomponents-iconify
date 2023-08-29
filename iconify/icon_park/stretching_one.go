package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StretchingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="24" cy="8" r="4" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M41 8L29 17.5909V44"/><path stroke-linecap="round" stroke-linejoin="round" d="M10.1111 23.25L19 18V28.9167L7 41"/></g>`),
		g.Group(children),
	)
}