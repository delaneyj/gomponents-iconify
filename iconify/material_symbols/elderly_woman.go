package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElderlyWoman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 23l-1.6-1.2L8.5 19H6q0-.875.488-3.013t1.337-4.362q.85-2.25 2.013-3.938T12.35 6q.925 0 1.288.575T14.675 8.3q.8 1.35 1.45 2.025t1.4 1.025q.275-.2.475-.275t.475-.075q.625 0 1.075.45T20 12.5V23h-1V12.5q0-.2-.15-.35T18.5 12q-.2 0-.35.15t-.15.35v1.25h-1v-.475q-.95-.525-1.95-1.363t-1.475-1.837L13 13l2 5.975V23h-2v-4h-2l-3 4Zm5.5-17.5q-.825 0-1.413-.588T11.5 3.5q0-.2.1-.6q-.275-.125-.438-.363T11 2q0-.425.288-.713T12 1q.3 0 .537.163t.363.437q.15-.05.3-.075t.3-.025q.825 0 1.413.588T15.5 3.5q0 .825-.588 1.413T13.5 5.5Z"/>`),
		g.Group(children),
	)
}