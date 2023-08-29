package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropThreeTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 18q-.825 0-1.413-.588T3 16V8q0-.825.588-1.413T5 6h14q.825 0 1.413.588T21 8v8q0 .825-.588 1.413T19 18H5Zm0-2h14V8H5v8Zm0 0V8v8Z"/>`),
		g.Group(children),
	)
}