package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sliders(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15q-1.25 0-2.125-.875T2 12q0-1.25.875-2.125T5 9h14q1.25 0 2.125.875T22 12q0 1.25-.875 2.125T19 15H5Zm9-2h5q.425 0 .713-.288T20 12q0-.425-.288-.713T19 11h-5v2Z"/>`),
		g.Group(children),
	)
}