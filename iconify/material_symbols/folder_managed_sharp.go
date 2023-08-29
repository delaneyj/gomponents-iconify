package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderManagedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 22l-.3-1.5q-.3-.125-.563-.263T15.6 19.9l-1.45.45l-1-1.7l1.15-1q-.05-.3-.05-.65t.05-.65l-1.15-1l1-1.7l1.45.45q.275-.2.538-.337t.562-.263L17 12h2l.3 1.5q.3.125.563.263t.537.337l1.45-.45l1 1.7l-1.15 1q.05.3.05.65t-.05.65l1.15 1l-1 1.7l-1.45-.45q-.275.2-.537.338t-.563.262L19 22h-2Zm1-3q.825 0 1.413-.588T20 17q0-.825-.588-1.413T18 15q-.825 0-1.413.588T16 17q0 .825.588 1.413T18 19ZM2 20V4h8l2 2h10v5.25q-.875-.625-1.9-.938T18 10q-2.925 0-4.963 2.038T11 17q0 .8.163 1.55t.512 1.45H2Z"/>`),
		g.Group(children),
	)
}