package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GestureLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 7.57a3 3 0 1 0 0 6H8v1.46a3 3 0 0 0 .409 1.511l2.409 4.13a3 3 0 0 0 3.54 1.334l6.09-2.03a3 3 0 0 0 2.052-2.846V9.055a3 3 0 0 0-1.17-2.378l-6.288-4.836l-1.527 1.017a2 2 0 0 0-.843 2.098l.581 2.614H4.5Zm-1 3a1 1 0 0 1 1-1h11.247l-1.122-5.048l.333-.222l5.152 3.963a1 1 0 0 1 .39.792v8.074a1 1 0 0 1-.684.949l-6.09 2.03a1 1 0 0 1-1.18-.445l-2.41-4.13A1 1 0 0 1 10 15.03v-3.46H4.5a1 1 0 0 1-1-1Z"/>`),
		g.Group(children),
	)
}