package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireTruck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21q-1.25 0-2.125-.875T4 18H3q-.825 0-1.413-.588T1 16v-5h11V7q0-.825.588-1.413T14 5h2V4q0-.425.288-.713T17 3h1q.425 0 .713.288T19 4v1h.55q.65 0 1.175.375t.725 1l1.45 4.3q.05.15.075.313t.025.337V18h-3q0 1.25-.875 2.125T17 21q-1.25 0-2.125-.875T14 18h-4q0 1.25-.875 2.125T7 21Zm0-2q.425 0 .713-.288T8 18q0-.425-.288-.713T7 17q-.425 0-.713.288T6 18q0 .425.288.713T7 19Zm10 0q.425 0 .713-.288T18 18q0-.425-.288-.713T17 17q-.425 0-.713.288T16 18q0 .425.288.713T17 19Zm-3-8h6.9l-1.35-4H14v4ZM1 10V8.5h1v-2H1V5h10v1.5h-1v2h1V10H1Zm2.5-1.5h1.75v-2H3.5v2Zm3.25 0H8.5v-2H6.75v2Z"/>`),
		g.Group(children),
	)
}