package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 23v-2H5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h5V1h2v22h-2Zm-5-5h5v-6l-5 6Zm9 3v-9l5 6V5h-5V3h5q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21h-5Z"/>`),
		g.Group(children),
	)
}