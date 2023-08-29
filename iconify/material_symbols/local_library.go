package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalLibrary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22.5q-1.8-1.7-4.125-2.6T3 19V8q2.525 0 4.85.913T12 11.55q1.825-1.725 4.15-2.638T21 8v11q-2.575 0-4.888.9T12 22.5ZM12 9q-1.65 0-2.825-1.175T8 5q0-1.65 1.175-2.825T12 1q1.65 0 2.825 1.175T16 5q0 1.65-1.175 2.825T12 9Z"/>`),
		g.Group(children),
	)
}