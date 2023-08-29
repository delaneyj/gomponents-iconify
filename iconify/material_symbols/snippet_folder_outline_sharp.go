package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnippetFolderOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h8l2 2h10v14H2Zm2-2h16V8h-8.825l-2-2H4v12Zm0 0V6v12Zm10.5-2.5v-5h1.375l1.625 1.625V15.5h-3ZM13 17h6v-5.5L16.5 9H13v8Z"/>`),
		g.Group(children),
	)
}