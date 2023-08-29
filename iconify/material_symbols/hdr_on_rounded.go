package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdrOnRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.75 15q-.325 0-.537-.213T16 14.25V10q0-.425.288-.713T17 9h2.5q.6 0 1.05.45T21 10.5v1q0 .575-.263.888t-.637.512l.475 1.125q.175.425-.012.7t-.638.275q-.175 0-.35-.113t-.25-.262L18.6 13h-1.1v1.25q0 .325-.213.537T16.75 15Zm.75-3.5h2v-1h-2v1ZM3.75 15q-.325 0-.537-.213T3 14.25v-4.5q0-.325.213-.537T3.75 9q.325 0 .537.213t.213.537V11h2V9.75q0-.325.213-.537T7.25 9q.325 0 .537.213T8 9.75v4.5q0 .325-.213.537T7.25 15q-.325 0-.537-.213T6.5 14.25V12.5h-2v1.75q0 .325-.213.537T3.75 15ZM10 15q-.2 0-.35-.15t-.15-.35v-5q0-.2.15-.35T10 9h3q.6 0 1.05.45t.45 1.05v3q0 .6-.45 1.05T13 15h-3Zm1-1.5h2v-3h-2v3Z"/>`),
		g.Group(children),
	)
}