package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14h12v-2H6v2Zm0-3h12V9H6v2Zm0-3h12V6H6v2ZM2 18V2h20v20l-4-4H2Z"/>`),
		g.Group(children),
	)
}