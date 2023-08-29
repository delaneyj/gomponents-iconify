package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M15.5 40.5h10c2.529 0 3-.529 3-3v-15h12l-20-20l-20 20h12v15c0 2.439.5 3 3 3z"/>`),
		g.Group(children),
	)
}