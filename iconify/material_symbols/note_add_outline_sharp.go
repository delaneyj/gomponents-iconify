package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteAddOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 18h2v-3h3v-2h-3v-3h-2v3H8v2h3v3Zm-7 4V2h10l6 6v14H4Zm9-13V4H6v16h12V9h-5ZM6 4v5v-5v16V4Z"/>`),
		g.Group(children),
	)
}