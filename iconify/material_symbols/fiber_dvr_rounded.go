package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiberDvrRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 15h3q.65 0 1.075-.425T9 13.5v-3q0-.65-.425-1.075T7.5 9h-3q-.2 0-.35.15T4 9.5v5q0 .2.15.35t.35.15Zm1-1.5v-3h2v3h-2Zm6.35-1.05L11 9.525q-.075-.225-.262-.375T10.3 9q-.35 0-.562.287t-.113.638l1.325 4.55q.075.225.263.375t.437.15h.4q.25 0 .437-.15t.263-.375l1.325-4.55q.1-.35-.113-.638T13.4 9q-.25 0-.438.15t-.262.375l-.85 2.925Zm4.65.55h1.15l.675 1.575q.075.2.25.313t.375.112q.375 0 .588-.3t.062-.65l-.5-1.15q.375-.2.637-.575T20 11.5v-1q0-.65-.425-1.075T18.5 9h-3q-.2 0-.35.15T15 9.5v4.75q0 .325.213.537t.537.213q.325 0 .537-.213t.213-.537V13Zm0-1.5v-1h2v1h-2ZM3 20q-.825 0-1.413-.588T1 18V6q0-.825.588-1.413T3 4h18q.825 0 1.413.588T23 6v12q0 .825-.588 1.413T21 20H3Z"/>`),
		g.Group(children),
	)
}