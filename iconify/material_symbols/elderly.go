package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elderly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 23l-1.6-1.2L9 18.325V13q0-.775.125-1.688T9.5 9.626L8 10.45V14H6V9.3l4.4-2.5q.625-.35 1.088-.537t.862-.188q.625 0 1.138.537T14.675 8.3q.8 1.35 1.45 2.025t1.4 1.025q.275-.2.475-.275t.475-.075q.625 0 1.075.45T20 12.5V23h-1V12.5q0-.2-.15-.35T18.5 12q-.2 0-.35.15t-.15.35v1.25h-1v-.475q-1.35-.575-2.1-1.288t-1.325-1.912q-.275.7-.437 1.713t-.113 1.912L15 16.5V23h-2v-5l-1.775-2.55L11 19l-3 4Zm5.5-17.5q-.825 0-1.413-.588T11.5 3.5q0-.825.588-1.413T13.5 1.5q.825 0 1.413.588T15.5 3.5q0 .825-.588 1.413T13.5 5.5Z"/>`),
		g.Group(children),
	)
}