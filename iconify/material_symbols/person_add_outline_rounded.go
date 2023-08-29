package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonAddOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 14q-.425 0-.713-.288T18 13v-2h-2q-.425 0-.713-.288T15 10q0-.425.288-.713T16 9h2V7q0-.425.288-.713T19 6q.425 0 .713.288T20 7v2h2q.425 0 .713.288T23 10q0 .425-.288.713T22 11h-2v2q0 .425-.288.713T19 14ZM9 12q-1.65 0-2.825-1.175T5 8q0-1.65 1.175-2.825T9 4q1.65 0 2.825 1.175T13 8q0 1.65-1.175 2.825T9 12Zm-7 8q-.425 0-.713-.288T1 19v-1.8q0-.85.438-1.563T2.6 14.55q1.55-.775 3.15-1.163T9 13q1.65 0 3.25.388t3.15 1.162q.725.375 1.163 1.088T17 17.2V19q0 .425-.288.713T16 20H2Zm1-2h12v-.8q0-.275-.138-.5t-.362-.35q-1.35-.675-2.725-1.012T9 15q-1.4 0-2.775.338T3.5 16.35q-.225.125-.363.35T3 17.2v.8Zm6-8q.825 0 1.413-.588T11 8q0-.825-.588-1.413T9 6q-.825 0-1.413.588T7 8q0 .825.588 1.413T9 10Zm0-2Zm0 7Z"/>`),
		g.Group(children),
	)
}