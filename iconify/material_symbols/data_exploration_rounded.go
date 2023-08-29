package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataExplorationRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 20.5q.425 0 .713-.288t.287-.712q0-.425-.288-.713T19.5 18.5q-.425 0-.713.288t-.287.712q0 .425.288.713t.712.287ZM12 22q-2.825 0-4.95-1.288T3.725 17.6l4.4-4.4l2.575 2.2q.15.125.325.188t.375.037q.175 0 .35-.075t.3-.2L16 11.425V12q0 .425.288.713T17 13q.425 0 .713-.288T18 12V9q0-.425-.288-.713T17 8h-3q-.425 0-.713.288T13 9q0 .425.288.713T14 10h.575L11.3 13.275L8.7 11.1q-.15-.125-.325-.188t-.35-.062q-.2.025-.375.1t-.3.2l-4.625 4.625q-.35-.875-.537-1.825T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12v8q0 .825-.588 1.413T20 22h-8Z"/>`),
		g.Group(children),
	)
}