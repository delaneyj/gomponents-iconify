package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChoppingBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M14 8C16.5033 8 35.7176 8 42.0112 8C43.1158 8 44 8.89543 44 10V38C44 39.1046 43.117 40 42.0125 40C35.8575 40 17.3255 40 14 40C10 40 4 38 4 24C4 10 11 8 14 8Z"/><path stroke="#fff" d="M12 20V28"/></g>`),
		g.Group(children),
	)
}