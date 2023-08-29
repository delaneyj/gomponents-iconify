package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Everdo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.07 37.76a2.25 2.25 0 0 1-3.17 0L4.5 26.37l3.17-3.17l9.81 9.8l22.85-22.84l3.17 3.17Z"/>`),
		g.Group(children),
	)
}