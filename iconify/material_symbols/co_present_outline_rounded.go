package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoPresentOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21V5H3v7q0 .425-.288.713T2 13q-.425 0-.713-.288T1 12V5q0-.825.588-1.413T3 3h18q.825 0 1.413.588T23 5v14q0 .825-.588 1.413T21 21ZM9 14q-1.65 0-2.825-1.175T5 10q0-1.65 1.175-2.825T9 6q1.65 0 2.825 1.175T13 10q0 1.65-1.175 2.825T9 14Zm0-2q.825 0 1.413-.588T11 10q0-.825-.588-1.413T9 8q-.825 0-1.413.588T7 10q0 .825.588 1.413T9 12ZM3 22q-.825 0-1.413-.588T1 20v-.8q0-.85.438-1.563T2.6 16.55q1.55-.775 3.15-1.163T9 15q1.65 0 3.25.388t3.15 1.162q.725.375 1.163 1.088T17 19.2v.8q0 .825-.588 1.413T15 22H3Zm0-2h12v-.8q0-.275-.138-.5t-.362-.35q-1.35-.675-2.725-1.012T9 17q-1.4 0-2.775.338T3.5 18.35q-.225.125-.363.35T3 19.2v.8Zm6-10Zm0 10Z"/>`),
		g.Group(children),
	)
}