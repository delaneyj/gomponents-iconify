package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhpRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.25 15q-.325 0-.537-.213T9.5 14.25v-4.5q0-.325.213-.537T10.25 9q.325 0 .537.213T11 9.75V11h2V9.75q0-.325.213-.537T13.75 9q.325 0 .537.213t.213.537v4.5q0 .325-.213.537T13.75 15q-.325 0-.537-.213T13 14.25V12.5h-2v1.75q0 .325-.213.537T10.25 15Zm-6.5 0q-.325 0-.537-.213T3 14.25V10q0-.425.288-.713T4 9h2.5q.6 0 1.05.45T8 10.5v1q0 .6-.45 1.05T6.5 13h-2v1.25q0 .325-.213.537T3.75 15Zm.75-3.5h2v-1h-2v1ZM17.25 15q-.325 0-.537-.213t-.213-.537V10q0-.425.288-.713T17.5 9H20q.6 0 1.05.45t.45 1.05v1q0 .6-.45 1.05T20 13h-2v1.25q0 .325-.213.537T17.25 15Zm.75-3.5h2v-1h-2v1Z"/>`),
		g.Group(children),
	)
}