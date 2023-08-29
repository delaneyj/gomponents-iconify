package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCopyOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19V1h9l6 6v12H6Zm8-11V3H8v14h11V8h-5ZM2 23V7h2v14h11v2H2ZM8 3v5v-5v14V3Z"/>`),
		g.Group(children),
	)
}