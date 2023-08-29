package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataServer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M41 4H7C5.34315 4 4 5.34315 4 7V41C4 42.6569 5.34315 44 7 44H41C42.6569 44 44 42.6569 44 41V7C44 5.34315 42.6569 4 41 4Z"/><path stroke="#fff" d="M4 32H44"/><path stroke="#fff" stroke-linejoin="round" d="M10 38H11"/><path stroke="#fff" stroke-linejoin="round" d="M26 38H38"/><path stroke="#000" stroke-linejoin="round" d="M44 37V27"/><path stroke="#000" stroke-linejoin="round" d="M4 37V27"/></g>`),
		g.Group(children),
	)
}