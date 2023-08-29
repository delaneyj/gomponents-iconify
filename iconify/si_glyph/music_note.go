package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.021 2.188v9.174a4.101 4.101 0 0 0-1.792.301c-1.605.649-2.533 2.066-2.074 3.162c.465 1.099 2.139 1.459 3.743.809c1.233-.5 1.958-1.45 2.067-2.362l-.007-8.885l7.062-1.359v6.378a4.089 4.089 0 0 0-1.944.297c-1.605.65-2.532 2.067-2.072 3.163c.463 1.098 2.137 1.459 3.742.809c1.233-.501 2.09-1.451 2.201-2.362L16.958.002L6.021 2.188z"/>`),
		g.Group(children),
	)
}