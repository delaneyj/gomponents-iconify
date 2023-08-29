package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M7.5 8.5v-8M2.618 2.499A6.963 6.963 0 0 0 .5 7.495c0 3.864 3.135 7.005 7 7.005c3.867 0 7-3.141 7-7.005A6.968 6.968 0 0 0 12.395 2.5"/>`),
		g.Group(children),
	)
}