package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M39.5 26.5v-10c0-2.529-.508-3-2.979-3H21.5v-12l-20 20l20 20v-12h15.021c2.44 0 2.979-.5 2.979-3z"/>`),
		g.Group(children),
	)
}