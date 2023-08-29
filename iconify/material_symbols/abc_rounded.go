package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbcRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 15q-.2 0-.35-.15t-.15-.35v-5q0-.2.15-.35T10 9h3.5q.425 0 .713.288T14.5 10v1q0 .425-.288.713T13.5 12q.425 0 .713.288T14.5 13v1q0 .425-.288.713T13.5 15H10Zm1-3.75h2v-.75h-2v.75Zm0 2.25h2v-.75h-2v.75ZM3.75 15q-.325 0-.537-.213T3 14.25V10q0-.425.288-.713T4 9h3q.425 0 .713.288T8 10v4.25q0 .325-.213.537T7.25 15q-.325 0-.537-.213T6.5 14.25v-.75h-2v.75q0 .325-.213.537T3.75 15Zm.75-3h2v-1.5h-2V12ZM17 15q-.425 0-.713-.288T16 14v-4q0-.425.288-.713T17 9h3q.425 0 .713.288T21 10v.5q0 .325-.213.537t-.537.213q-.325 0-.537-.213T19.5 10.5h-2v3h2q0-.325.213-.537t.537-.213q.325 0 .537.213T21 13.5v.5q0 .425-.288.713T20 15h-3Z"/>`),
		g.Group(children),
	)
}