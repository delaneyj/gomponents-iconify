package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscFullRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16.5q1.875 0 3.188-1.313T16.5 12q0-1.875-1.313-3.188T12 7.5q-1.875 0-3.188 1.313T7.5 12q0 1.875 1.313 3.188T12 16.5Zm0-3.5q-.425 0-.713-.288T11 12q0-.425.288-.713T12 11q.425 0 .713.288T13 12q0 .425-.288.713T12 13Zm9 5q-.425 0-.713-.288T20 17v-6q0-.425.288-.713T21 10q.425 0 .713.288T22 11v6q0 .425-.288.713T21 18Zm0 4q-.425 0-.713-.288T20 21q0-.425.288-.713T21 20q.425 0 .713.288T22 21q0 .425-.288.713T21 22Zm-9 0q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q3 0 5.475 1.638T21.15 8H19q-.425 0-.712.288T18 9v11q-1.325.95-2.85 1.475T12 22Z"/>`),
		g.Group(children),
	)
}