package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCopyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 19H8q-.825 0-1.413-.588T6 17V3q0-.825.588-1.413T8 1h7l6 6v10q0 .825-.588 1.413T19 19ZM14 8V3H8v14h11V8h-5ZM4 23q-.825 0-1.413-.588T2 21V7h2v14h11v2H4ZM8 3v5v-5v14V3Z"/>`),
		g.Group(children),
	)
}