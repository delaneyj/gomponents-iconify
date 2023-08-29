package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropFree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19v-4h2v4h4v2H5Zm10 0v-2h4v-4h2v4q0 .825-.588 1.413T19 21h-4ZM3 9V5q0-.825.588-1.413T5 3h4v2H5v4H3Zm16 0V5h-4V3h4q.825 0 1.413.588T21 5v4h-2Z"/>`),
		g.Group(children),
	)
}