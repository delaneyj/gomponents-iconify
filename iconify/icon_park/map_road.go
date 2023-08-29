package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapRoad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M41 4H7C5.34315 4 4 5.34315 4 7V41C4 42.6569 5.34315 44 7 44H41C42.6569 44 44 42.6569 44 41V7C44 5.34315 42.6569 4 41 4Z"/><path stroke="#fff" stroke-linecap="round" d="M33 12L37 36"/><path stroke="#fff" stroke-linecap="round" d="M16 12L12 36"/><path stroke="#fff" stroke-linecap="round" d="M24 12V16"/><path stroke="#fff" stroke-linecap="round" d="M24 22V26"/><path stroke="#fff" stroke-linecap="round" d="M24 32V36"/></g>`),
		g.Group(children),
	)
}