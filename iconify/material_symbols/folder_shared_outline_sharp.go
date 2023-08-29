package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSharedOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h8l2 2h10v14H2Zm2-2h16V8h-8.825l-2-2H4v12Zm0 0V6v12Zm7-1h8v-.55q0-1.125-1.1-1.788T15 14q-1.8 0-2.9.663T11 16.45V17Zm4-4q.825 0 1.413-.588T17 11q0-.825-.588-1.413T15 9q-.825 0-1.413.588T13 11q0 .825.588 1.413T15 13Z"/>`),
		g.Group(children),
	)
}