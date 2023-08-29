package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSupervisedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 20v-.55q0-1.125 1.1-1.788T18 17q1.8 0 2.9.663T22 19.45V20h-8Zm4-4q-.825 0-1.413-.588T16 14q0-.825.588-1.413T18 12q.825 0 1.413.588T20 14q0 .825-.588 1.413T18 16ZM2 20V4h8l2 2h10v5.25q-.875-.625-1.9-.938T18 10q-2.925 0-4.963 2.038T11 17q0 .8.163 1.55t.512 1.45H2Z"/>`),
		g.Group(children),
	)
}