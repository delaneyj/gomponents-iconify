package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpoTwoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-.425 0-.713-.288T11 19v-4q0-.425.288-.713T12 14h2.5q.425 0 .713.288T15.5 15v4q0 .425-.288.713T14.5 20H12Zm.5-1.5H14v-3h-1.5v3ZM18 22q-.425 0-.713-.288T17 21v-1.75q0-.425.288-.713T18 18.25h2v-.75h-2.25q-.325 0-.537-.213T17 16.75q0-.325.213-.537T17.75 16h2.75q.425 0 .713.288T21.5 17v1.75q0 .425-.288.713t-.712.287h-2v.75h2.25q.325 0 .537.213t.213.537q0 .325-.213.537T20.75 22H18Zm-9-.05q-3.075-.35-5.038-2.638T2 13.8q0-1.625.725-3.213T4.55 7.564q1.1-1.438 2.388-2.713T9.325 2.6q.15-.125.313-.187T10 2.35q.2 0 .363.063t.312.187q1.025.9 2.125 1.975t2.075 2.263q.975 1.187 1.738 2.487T17.75 12H11q-.825 0-1.413.588T9 14v7.95Z"/>`),
		g.Group(children),
	)
}