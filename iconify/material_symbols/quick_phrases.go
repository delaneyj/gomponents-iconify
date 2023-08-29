package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickPhrases(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22.5v-5.525q-2.525-.2-4.263-2.05T2 10.5q0-2.725 1.888-4.612T8.5 4h.675L7.6 2.4L9 1l4 4l-4 4l-1.4-1.4L9.175 6H8.5Q6.625 6 5.312 7.313T4 10.5q0 1.875 1.313 3.188T8.5 15H10v2.675L12.675 15H15.5q1.875 0 3.188-1.313T20 10.5q0-1.875-1.313-3.188T15.5 6H15V4h.5q2.725 0 4.612 1.888T22 10.5q0 2.725-1.888 4.612T15.5 17h-2L8 22.5Z"/>`),
		g.Group(children),
	)
}