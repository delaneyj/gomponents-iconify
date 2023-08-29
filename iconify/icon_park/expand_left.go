package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M6 9C6 7.34315 7.34315 6 9 6H39C40.6569 6 42 7.34315 42 9V39C42 40.6569 40.6569 42 39 42H9C7.34315 42 6 40.6569 6 39V9Z"/><path stroke="#fff" stroke-linecap="round" d="M32 6V42"/><path stroke="#fff" stroke-linecap="round" d="M16 20L20 24L16 28"/><path stroke="#000" stroke-linecap="round" d="M26 6H38"/><path stroke="#000" stroke-linecap="round" d="M26 42H38"/></g>`),
		g.Group(children),
	)
}