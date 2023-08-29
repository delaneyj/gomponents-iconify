package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h2.1l6-5.95l-2.15-2.15L7 14.85V17Zm8.8-6.65l1.05-1.1Q17 9.1 17 8.9t-.15-.35l-1.4-1.4Q15.3 7 15.1 7t-.35.15l-1.1 1.05l2.15 2.15ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h4.2q.325-.9 1.088-1.45T12 1q.95 0 1.713.55T14.8 3H19q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm7-16.75q.325 0 .537-.213t.213-.537q0-.325-.213-.537T12 2.75q-.325 0-.537.213t-.213.537q0 .325.213.537T12 4.25Z"/>`),
		g.Group(children),
	)
}