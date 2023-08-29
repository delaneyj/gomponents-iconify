package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="10" r="6" fill="#2F88FF"/><path fill="#2F88FF" d="M30 36H18L10 16H38L30 36Z"/><path d="M27 36V44"/><path d="M21 36V44"/></g>`),
		g.Group(children),
	)
}