package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResetImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 9V3h2v3.35Q6.25 4.8 8.063 3.9T12 3q2.95 0 5.263 1.675T20.5 9h-2.175q-.85-1.8-2.525-2.9T12 5q-1.425 0-2.688.525T7.1 7H9v2H3Zm3 9h12l-3.75-5l-3 4L9 14l-3 4Zm-1 4q-.825 0-1.413-.588T3 20v-8h2v8h14v-8h2v8q0 .825-.588 1.413T19 22H5Z"/>`),
		g.Group(children),
	)
}