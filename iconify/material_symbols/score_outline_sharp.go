package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScoreOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm6-8.5l4 4l6-6V5H5v11.5l4-4Zm3-.5V6h1.5v6H12Zm3.5 0l-2-3l2-3h1.7l-2 3l2 3h-1.7ZM7 12V8.25h2.5V7.5H7V6h4v3.75H8.5v.75H11V12H7Zm2 3l-4 4h14v-6l-6 6l-4-4Zm-4 4V5v14Z"/>`),
		g.Group(children),
	)
}