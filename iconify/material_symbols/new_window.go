package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewWindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h6v2H5v14h14v-6h2v6q0 .825-.588 1.413T19 21H5Zm11-10V8h-3V6h3V3h2v3h3v2h-3v3h-2Z"/>`),
		g.Group(children),
	)
}