package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Editing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="13" cy="35" r="7" fill="#2F88FF"/><circle cx="35" cy="35" r="7" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 6L32 28"/><path stroke-linecap="round" stroke-linejoin="round" d="M40 6L16 28"/></g>`),
		g.Group(children),
	)
}