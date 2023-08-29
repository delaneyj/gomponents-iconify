package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConversionPath(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 21q-.975 0-1.75-.563T16.175 19H11q-1.65 0-2.825-1.175T7 15q0-1.65 1.175-2.825T11 11h2q.825 0 1.413-.588T15 9q0-.825-.588-1.413T13 7H7.825q-.325.875-1.088 1.438T5 9q-1.25 0-2.125-.875T2 6q0-1.25.875-2.125T5 3q.975 0 1.738.563T7.824 5H13q1.65 0 2.825 1.175T17 9q0 1.65-1.175 2.825T13 13h-2q-.825 0-1.413.588T9 15q0 .825.588 1.413T11 17h5.175q.325-.875 1.088-1.438T19 15q1.25 0 2.125.875T22 18q0 1.25-.875 2.125T19 21ZM5 7q.425 0 .713-.288T6 6q0-.425-.288-.713T5 5q-.425 0-.713.288T4 6q0 .425.288.713T5 7Z"/>`),
		g.Group(children),
	)
}