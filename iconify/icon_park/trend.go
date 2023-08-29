package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M39 6H9C7.34315 6 6 7.34315 6 9V39C6 40.6569 7.34315 42 9 42H39C40.6569 42 42 40.6569 42 39V9C42 7.34315 40.6569 6 39 6Z"/><path stroke="#fff" stroke-linecap="round" d="M13.4398 29.8347L19.0967 24.1778L23.4847 28.5555L34 18.0001"/><path stroke="#fff" stroke-linecap="round" d="M26 18H34V26"/></g>`),
		g.Group(children),
	)
}