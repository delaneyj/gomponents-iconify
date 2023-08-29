package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenshotMonitorRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 14.5h-1.75q-.325 0-.537.213T15 15.25q0 .325.213.537t.537.213H18q.425 0 .713-.288T19 15v-2.25q0-.325-.213-.537T18.25 12q-.325 0-.537.213t-.213.537v1.75Zm-11-7h1.75q.325 0 .537-.213T9 6.75q0-.325-.213-.537T8.25 6H6q-.425 0-.713.288T5 7v2.25q0 .325.213.537T5.75 10q.325 0 .537-.213T6.5 9.25V7.5ZM4 19q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v12q0 .825-.588 1.413T20 19h-4v1q0 .425-.288.713T15 21H9q-.425 0-.713-.288T8 20v-1H4Z"/>`),
		g.Group(children),
	)
}