package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropFiveFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20q-.825 0-1.413-.588T3 18V6q0-.825.588-1.413T5 4h14q.825 0 1.413.588T21 6v12q0 .825-.588 1.413T19 20H5Z"/>`),
		g.Group(children),
	)
}