package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h6l2 2h8q.825 0 1.413.588T22 8v10q0 .825-.588 1.413T20 20H4Zm10.5-3h2q.625 0 1.063-.438T18 15.5v-4h1V10h-2.5V9h-2v1H12v1.5h1v4q0 .625.438 1.063T14.5 17Zm0-5.5h2v4h-2v-4Z"/>`),
		g.Group(children),
	)
}