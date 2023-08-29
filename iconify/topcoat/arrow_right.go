package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M2.5 26.5v-10c0-2.529.529-3 3-3h15v-12l20 20l-20 20v-12h-15c-2.439 0-3-.5-3-3z"/>`),
		g.Group(children),
	)
}