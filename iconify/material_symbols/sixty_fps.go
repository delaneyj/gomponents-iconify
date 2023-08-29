package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixtyFps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19H5q-1.25 0-2.125-.875T2 16V8q0-1.25.875-2.125T5 5h5v3H5v2h3q1.25 0 2.125.875T11 13v3q0 1.25-.875 2.125T8 19Zm-3-6v3h3v-3H5Zm10 3h4V8h-4v8Zm0 3q-1.25 0-2.125-.875T12 16V8q0-1.25.875-2.125T15 5h4q1.25 0 2.125.875T22 8v8q0 1.25-.875 2.125T19 19h-4Z"/>`),
		g.Group(children),
	)
}