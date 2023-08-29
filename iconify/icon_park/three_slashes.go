package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeSlashes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 10L42 4V11L6 17V10Z"/><path d="M6 24L42 18V25L6 31V24Z"/><path d="M6 38L42 32V39L6 45V38Z"/></g>`),
		g.Group(children),
	)
}