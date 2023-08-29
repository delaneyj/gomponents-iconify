package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkdownCopyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 18q-.825 0-1.413-.588T7 16V4q0-.825.588-1.413T9 2h9q.825 0 1.413.588T20 4v12q0 .825-.588 1.413T18 18H9Zm-4 4q-.825 0-1.413-.588T3 20V7q0-.425.288-.713T4 6q.425 0 .713.288T5 7v13h10q.425 0 .713.288T16 21q0 .425-.288.713T15 22H5Zm6.75-13.5h1v2.25q0 .325.213.537t.537.213q.325 0 .537-.213t.213-.537V8.5h1v3.75q0 .325.213.537T16 13q.325 0 .537-.213t.213-.537V8q0-.425-.288-.713T15.75 7h-4.5q-.425 0-.713.288T10.25 8v4.25q0 .325.213.537T11 13q.325 0 .537-.213t.213-.537V8.5Z"/>`),
		g.Group(children),
	)
}