package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneHundredTwentyThreeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 10.5h-.75q-.325 0-.537-.213T4 9.75q0-.325.213-.537T4.75 9h1.5q.325 0 .537.213T7 9.75v4.5q0 .325-.213.537T6.25 15q-.325 0-.537-.213T5.5 14.25V10.5ZM9 14.25V12.5q0-.425.288-.713T10 11.5h2v-1H9.75q-.325 0-.537-.213T9 9.75q0-.325.213-.537T9.75 9h2.75q.425 0 .713.288T13.5 10v1.5q0 .425-.288.713t-.712.287h-2v1h2.25q.325 0 .537.213t.213.537q0 .325-.213.537T12.75 15h-3q-.325 0-.537-.213T9 14.25Zm9.5.75h-2.75q-.325 0-.537-.213T15 14.25q0-.325.213-.537t.537-.213H18v-1h-1.5q-.2 0-.35-.15T16 12q0-.2.15-.35t.35-.15H18v-1h-2.25q-.325 0-.537-.213T15 9.75q0-.325.213-.537T15.75 9h2.75q.425 0 .713.288T19.5 10v4q0 .425-.288.713T18.5 15Z"/>`),
		g.Group(children),
	)
}