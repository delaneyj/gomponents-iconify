package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Editor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M40 33V42C40 43.1046 39.1046 44 38 44H31.5"/><path stroke-linejoin="round" d="M40 16V6C40 4.89543 39.1046 4 38 4H10C8.89543 4 8 4.89543 8 6V42C8 43.1046 8.89543 44 10 44H16"/><path d="M16 16H30"/><path d="M23 44L40 23"/><path d="M16 24H24"/></g>`),
		g.Group(children),
	)
}