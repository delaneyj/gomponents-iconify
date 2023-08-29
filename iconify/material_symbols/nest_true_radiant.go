package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestTrueRadiant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-1.25 0-2.125-.875T3 18V8h2v10q0 .425.288.713T6 19q.425 0 .713-.288T7 18V6q0-1.25.875-2.125T10 3q1.25 0 2.125.875T13 6v12q0 .425.288.713T14 19q.425 0 .713-.288T15 18V6q0-1.25.875-2.125T18 3q1.25 0 2.125.875T21 6v10h-2V6q0-.425-.288-.713T18 5q-.425 0-.713.288T17 6v12q0 1.25-.875 2.125T14 21q-1.25 0-2.125-.875T11 18V6q0-.425-.288-.713T10 5q-.425 0-.713.288T9 6v12q0 1.25-.875 2.125T6 21Z"/>`),
		g.Group(children),
	)
}