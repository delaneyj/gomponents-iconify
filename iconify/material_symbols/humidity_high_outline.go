package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HumidityHighOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-3.325 0-5.663-2.313T4 13.1q0-1.65.625-3.05t1.725-2.5L12 2l5.65 5.55q1.1 1.1 1.725 2.5T20 13.1q0 3.275-2.337 5.588T12 21Z"/>`),
		g.Group(children),
	)
}