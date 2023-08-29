package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GuideBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 4V41"/><path fill="#2F88FF" d="M24 8H39.5455L42 12L39.5455 16H24V8Z"/><path fill="#2F88FF" d="M24 22H8.45455L6 26L8.45455 30H24V22Z"/><path stroke-linecap="round" d="M16 42H32"/></g>`),
		g.Group(children),
	)
}