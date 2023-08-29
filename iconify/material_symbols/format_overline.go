package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatOverline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5V3h14v2H5Zm7 16q-2.925 0-4.963-2.038T5 14q0-2.925 2.038-4.963T12 7q2.925 0 4.963 2.038T19 14q0 2.925-2.038 4.963T12 21Zm0-2.5q1.875 0 3.188-1.313T16.5 14q0-1.875-1.313-3.188T12 9.5q-1.875 0-3.188 1.313T7.5 14q0 1.875 1.313 3.188T12 18.5Z"/>`),
		g.Group(children),
	)
}