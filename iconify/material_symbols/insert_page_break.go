package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertPageBreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20v-3h16v3q0 .825-.588 1.413T18 22H6Zm7-13h5l-5-5v5Zm-4 6v-2h6v2H9Zm8 0v-2h6v2h-6ZM1 15v-2h6v2H1Zm3-4V4q0-.825.588-1.413T6 2h8l6 6v3H4Z"/>`),
		g.Group(children),
	)
}