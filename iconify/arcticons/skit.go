package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 15.767L32.885 5.8l4.124 11.333l-27.382 9.966zm5.494 15.098L38.377 20.9L42.5 32.23l-27.38 9.968zM30.16 19.62l8.22 1.29M9.62 27.09l8.22 1.29"/>`),
		g.Group(children),
	)
}