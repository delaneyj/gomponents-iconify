package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M18 24H31"/><path d="M4 17C4 15.8954 4.89543 15 6 15H42C43.1046 15 44 15.8954 44 17V31C44 32.1046 43.1046 33 42 33H6C4.89543 33 4 32.1046 4 31V17Z"/></g>`),
		g.Group(children),
	)
}