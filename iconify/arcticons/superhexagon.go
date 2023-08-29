package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Superhexagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.22 21l-3.01 11.21L21 35.22L12.79 27l3-11.21l11.21-3l2.49-9.29l-20.5 5.49l-5.49 20.5L18.51 44.5l20.5-5.49l5.49-20.5L35.22 21z"/>`),
		g.Group(children),
	)
}