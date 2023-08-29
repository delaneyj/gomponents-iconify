package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridOnOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h3.325v-3.325H5V19Zm5.325 0h3.35v-3.325h-3.35V19Zm5.35 0H19v-3.325h-3.325V19ZM5 13.675h3.325v-3.35H5v3.35Zm5.325 0h3.35v-3.35h-3.35v3.35Zm5.35 0H19v-3.35h-3.325v3.35ZM5 8.325h3.325V5H5v3.325Zm5.325 0h3.35V5h-3.35v3.325Zm5.35 0H19V5h-3.325v3.325Z"/>`),
		g.Group(children),
	)
}