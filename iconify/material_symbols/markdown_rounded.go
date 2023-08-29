package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkdownRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.25 12.125l-.675-.675q-.225-.225-.525-.213t-.525.238Q13.3 11.7 13.3 12t.225.525L15.3 14.3q.3.3.7.3t.7-.3l1.775-1.775Q18.7 12.3 18.7 12t-.225-.525q-.225-.225-.538-.225t-.537.225l-.65.65V9.75q0-.325-.213-.537T16 9q-.325 0-.537.213t-.213.537v2.375ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm3-9.5h1v2.25q0 .325.213.537t.537.213q.325 0 .537-.213t.213-.537V10.5h1v3.75q0 .325.213.537t.537.213q.325 0 .537-.213T12 14.25V10q0-.425-.288-.713T11 9H6.5q-.425 0-.713.288T5.5 10v4.25q0 .325.213.537T6.25 15q.325 0 .537-.213T7 14.25V10.5Z"/>`),
		g.Group(children),
	)
}