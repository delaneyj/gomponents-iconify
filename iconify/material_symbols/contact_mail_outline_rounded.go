package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactMailOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 11h5q.425 0 .713-.288T21 10V7q0-.425-.288-.713T20 6h-5q-.425 0-.713.288T14 7v3q0 .425.288.713T15 11Zm2.5-2.25l1.85-1.3q.2-.15.425-.038t.225.363q0 .025-.175.35l-1.75 1.225q-.275.2-.575.2t-.575-.2l-1.75-1.225Q15.15 8.1 15 7.775q0-.25.225-.363t.425.038l1.85 1.3ZM2 21q-.825 0-1.413-.588T0 19V5q0-.825.588-1.413T2 3h20q.825 0 1.413.588T24 5v14q0 .825-.588 1.413T22 21H2Zm13.9-2H22V5H2v14h.1q1.05-1.875 2.9-2.938T9 15q2.15 0 4 1.063T15.9 19ZM9 14q1.25 0 2.125-.875T12 11q0-1.25-.875-2.125T9 8q-1.25 0-2.125.875T6 11q0 1.25.875 2.125T9 14Zm-4.45 5h8.9q-.85-.95-2.013-1.475T9 17q-1.275 0-2.425.525T4.55 19ZM9 12q-.425 0-.713-.288T8 11q0-.425.288-.713T9 10q.425 0 .713.288T10 11q0 .425-.288.713T9 12Zm3 0Z"/>`),
		g.Group(children),
	)
}