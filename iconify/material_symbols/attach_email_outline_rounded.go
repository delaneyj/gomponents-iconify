package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttachEmailOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18q-.825 0-1.413-.588T1 16V4q0-.825.588-1.413T3 2h16q.825 0 1.413.588T21 4v5h-2V6l-7.475 4.675q-.125.075-.263.113t-.262.037q-.125 0-.263-.037t-.262-.113L3 6v10h10v2H3Zm8-9l8-5H3l8 5ZM3 6v.25v-1.475v.025V4v.8v-.025V6.25V6v10V6Zm16 16q-1.65 0-2.825-1.175T15 18v-4.5q0-1.05.725-1.775T17.5 11q1.05 0 1.775.725T20 13.5V17q0 .425-.288.713T19 18q-.425 0-.713-.288T18 17v-3.5q0-.2-.15-.35T17.5 13q-.2 0-.35.15t-.15.35V18q0 .825.588 1.413T19 20q.825 0 1.413-.588T21 18v-3q0-.425.288-.713T22 14q.425 0 .713.288T23 15v3q0 1.65-1.175 2.825T19 22Z"/>`),
		g.Group(children),
	)
}