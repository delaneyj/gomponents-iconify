package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#000" d="M27 24V37.125C27 39 26.6379 44 23 44C19.5714 44 19 39.625 19 38.375"/><path fill="#2F88FF" stroke="#000" d="M24 4C38.5 4 43.375 17.3333 44 24H4C4.625 17.3333 9.5 4 24 4Z"/><path stroke="#fff" d="M19 14L23 18L29 11"/></g>`),
		g.Group(children),
	)
}