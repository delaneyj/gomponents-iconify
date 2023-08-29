package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContrastView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="17" height="34" x="7" y="7" fill="#2F88FF"/><path stroke-linecap="round" d="M24 7H28"/><path stroke-linecap="round" d="M33 7H35"/><path stroke-linecap="round" d="M33 41H35"/><path stroke-linecap="round" d="M41 7V9"/><path stroke-linecap="round" d="M41 15V17"/><path stroke-linecap="round" d="M41 23V25"/><path stroke-linecap="round" d="M41 31V33"/><path stroke-linecap="round" d="M41 39V41"/><path stroke-linecap="round" d="M27 41H24"/><path stroke-linecap="round" d="M24 4V44"/></g>`),
		g.Group(children),
	)
}