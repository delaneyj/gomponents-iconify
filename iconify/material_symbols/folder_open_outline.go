package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h6l2 2h8q.825 0 1.413.588T22 8H11.175l-2-2H4v12l2.4-8h17.1l-2.575 8.575q-.2.65-.738 1.038T19 20H4Zm2.1-2H19l1.8-6H7.9l-1.8 6Zm0 0l1.8-6l-1.8 6ZM4 8V6v2Z"/>`),
		g.Group(children),
	)
}