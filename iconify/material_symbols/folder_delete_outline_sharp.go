package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderDeleteOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h8l2 2h10v14H2Zm2-2h16V8h-8.825l-2-2H4v12Zm0 0V6v12Zm9-1h5v-5.5h1V10h-2.5V9h-2v1H12v1.5h1V17Zm1.5-1.5v-4h2v4h-2Z"/>`),
		g.Group(children),
	)
}