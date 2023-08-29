package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSnippetOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19h14V9.825L14.175 5H5v14Zm0 2q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h10l6 6v10q0 .825-.588 1.413T19 21H5Zm2-4h10v-2H7v2Zm0-4h10v-2H7v2Zm0-4h7V7H7v2ZM5 19V5v14Z"/>`),
		g.Group(children),
	)
}