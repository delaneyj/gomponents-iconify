package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStartArrowOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.475 18.025l-8.15-5.175q-.475-.3-.475-.85t.475-.85l8.15-5.175q.5-.325 1.012-.037t.513.887V11h8q.425 0 .713.288T22 12q0 .425-.288.713T21 13h-8v4.175q0 .6-.513.888t-1.012-.038ZM11 15.35v-6.7L5.725 12L11 15.35ZM11 12Z"/>`),
		g.Group(children),
	)
}