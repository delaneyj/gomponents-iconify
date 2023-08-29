package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalActivitySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.2 16l2.8-2.1l2.75 2.1l-1.05-3.4l2.8-2.2h-3.4L12 7l-1.1 3.4H7.5l2.75 2.2L9.2 16ZM2 20v-6q.825 0 1.413-.588T4 12q0-.825-.588-1.413T2 10V4h20v6q-.825 0-1.413.588T20 12q0 .825.588 1.413T22 14v6H2Z"/>`),
		g.Group(children),
	)
}