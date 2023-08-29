package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Netx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.977 10.298l-19.43 15.259L6.762 10.405A19.4 19.4 0 0 1 13.469 5.5l11.62 15.258M9.023 37.702l19.43-15.258l12.785 15.151A19.4 19.4 0 0 1 34.53 42.5L22.91 27.242"/>`),
		g.Group(children),
	)
}