package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockResetSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21q-1.85 0-3.463-.688T6.7 18.425L8.125 17q.95.925 2.2 1.463T13 19q2.9 0 4.95-2.05T20 12q0-2.9-2.05-4.95T13 5q-2.9 0-4.95 2.05T6 12v.175l1.825-1.825l1.425 1.4L5 16L.75 11.75l1.425-1.4L4 12.2V12q0-1.875.713-3.513t1.925-2.85q1.212-1.212 2.85-1.924T13 3q1.875 0 3.513.713t2.85 1.924q1.212 1.213 1.925 2.85T22 12q0 3.75-2.625 6.375T13 21Zm-3-5v-5h1v-1q0-.825.588-1.413T13 8q.825 0 1.413.588T15 10v1h1v5h-6Zm2-5h2v-1q0-.425-.288-.713T13 9q-.425 0-.713.288T12 10v1Z"/>`),
		g.Group(children),
	)
}