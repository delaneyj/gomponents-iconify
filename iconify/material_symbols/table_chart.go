package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 8V5q0-.825.588-1.413T5 3h15q.825 0 1.413.588T22 5v3H3Zm5 2v11H5q-.825 0-1.413-.588T3 19v-9h5Zm9 0h5v9q0 .825-.588 1.413T20 21h-3V10Zm-2 0v11h-5V10h5Z"/>`),
		g.Group(children),
	)
}