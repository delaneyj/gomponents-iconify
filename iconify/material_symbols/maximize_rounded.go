package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaximizeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5q-.425 0-.713-.288T4 4q0-.425.288-.713T5 3h14q.425 0 .713.288T20 4q0 .425-.288.713T19 5H5Z"/>`),
		g.Group(children),
	)
}