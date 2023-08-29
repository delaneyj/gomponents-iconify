package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M20.5 5.5C9.454 5.5.5 12.887.5 22s8.954 16.5 20 16.5s20-7.387 20-16.5s-8.954-16.5-20-16.5z"/>`),
		g.Group(children),
	)
}