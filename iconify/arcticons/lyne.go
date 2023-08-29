package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lyne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.094 10.998l6.499 6.498l-6.499 6.499l-6.498-6.499zM24.098 11l6.498 6.499l-6.498 6.498l-6.499-6.498zm13.002-.007l6.497 6.498l-6.498 6.498l-6.498-6.498zM11.102 24l6.498 6.5l-6.498 6.498L4.603 30.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.094 23.992l6.499 6.499l-6.499 6.498l-6.498-6.498zm13.003.003l6.498 6.498l-6.498 6.498l-6.499-6.498z"/>`),
		g.Group(children),
	)
}