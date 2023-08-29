package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenShareRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 11h2v1.4q0 .175.15.238t.275-.063l2.225-2.225q.15-.15.15-.35t-.15-.35l-2.225-2.225q-.125-.125-.275-.062T13 7.6V9h-2q-1.25 0-2.125.875T8 12v1q0 .425.288.713T9 14q.425 0 .713-.288T10 13v-1q0-.425.288-.713T11 11ZM2 21q-.425 0-.713-.288T1 20q0-.425.288-.713T2 19h20q.425 0 .713.288T23 20q0 .425-.288.713T22 21H2Zm2-3q-.825 0-1.413-.588T2 16V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v11q0 .825-.588 1.413T20 18H4Z"/>`),
		g.Group(children),
	)
}