package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PostAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h9v2H5v14h14v-9h2v9q0 .825-.588 1.413T19 21H5Zm3-4v-2h8v2H8Zm0-3v-2h8v2H8Zm0-3V9h8v2H8Zm9-2V7h-2V5h2V3h2v2h2v2h-2v2h-2Z"/>`),
		g.Group(children),
	)
}