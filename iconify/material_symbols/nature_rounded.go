package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatureRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.425 0-.713-.288T5 21q0-.425.288-.713T6 20h5v-4H9q-2.075 0-3.538-1.463T4 11q0-1.5.825-2.763T7.05 6.4q.225-1.875 1.638-3.138T12 2q1.9 0 3.313 1.263T16.95 6.4q1.4.575 2.225 1.838T20 11q0 2.075-1.463 3.538T15 16h-2v4h5q.425 0 .713.288T19 21q0 .425-.288.713T18 22H6Z"/>`),
		g.Group(children),
	)
}