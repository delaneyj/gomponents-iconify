package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderZipOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 12v-2h2v2h-2Zm0 2h-2v-2h2v2Zm0 2v-2h2v2h-2Zm-4.825-8l-2-2H4v12h10v-2h2v2h4V8h-4v2h-2V8h-2.825ZM2 20V4h8l2 2h10v14H2Zm2-2V6v12Z"/>`),
		g.Group(children),
	)
}