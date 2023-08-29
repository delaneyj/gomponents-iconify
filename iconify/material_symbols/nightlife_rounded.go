package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NightlifeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 20q-.425 0-.713-.288T5.5 19q0-.425.288-.713T6.5 18h1v-4L2.525 6.55q-.35-.5-.05-1.025t.9-.525h10.25q.6 0 .9.525t-.05 1.025L9.5 14v4h1q.425 0 .713.288T11.5 19q0 .425-.288.713T10.5 20h-4ZM6.4 9h4.2L12 7H5l1.4 2Zm10.1 11q-1.25 0-2.125-.875T13.5 17q0-1.25.875-2.125T16.5 14q.275 0 .525.038t.475.162V7q0-.825.588-1.413T19.5 5H21q.625 0 1.063.438T22.5 6.5q0 .625-.438 1.063T21 8h-1.5v9q0 1.25-.875 2.125T16.5 20Z"/>`),
		g.Group(children),
	)
}