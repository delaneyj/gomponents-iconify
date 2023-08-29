package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStartArrowNotchRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.575 13L12 17.25q.2.35-.075.613t-.625.062L3.325 12.85q-.475-.3-.475-.85t.475-.85L11.3 6.075q.35-.2.625.063T12 6.75L9.575 11H21q.425 0 .713.288T22 12q0 .425-.288.713T21 13H9.575Z"/>`),
		g.Group(children),
	)
}