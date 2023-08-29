package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkdownCopyOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18V2h13v16H7Zm2-2h9V4H9v12Zm-6 6V6h2v14h11v2H3Zm7.25-9h1.5V8.5h1v3h1.5v-3h1V13h1.5V7h-6.5v6ZM9 16V4v12Z"/>`),
		g.Group(children),
	)
}