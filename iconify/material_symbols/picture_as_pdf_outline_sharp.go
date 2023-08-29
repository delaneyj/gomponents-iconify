package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureAsPdfOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 12.5h1v-2h1.5l.5-.5V8l-.5-.5H9v5Zm1-3v-1h1v1h-1Zm3 3h2.5l.5-.5V8l-.5-.5H13v5Zm1-1v-3h1v3h-1Zm3 1h1v-2h1v-1h-1v-1h1v-1h-2v5ZM6 18V2h16v16H6Zm2-2h12V4H8v12Zm-6 6V6h2v14h14v2H2ZM8 4v12V4Z"/>`),
		g.Group(children),
	)
}