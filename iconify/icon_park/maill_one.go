package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaillOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M39 6H9C7.34315 6 6 7.34315 6 9V39C6 40.6569 7.34315 42 9 42H39C40.6569 42 42 40.6569 42 39V9C42 7.34315 40.6569 6 39 6Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M36 16H12V32H36V16Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M36 16L24 24L12 16"/></g>`),
		g.Group(children),
	)
}