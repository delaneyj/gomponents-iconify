package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.6 3.5L12.1 7H0v2h12.1l-3.5 3.5l1.4 1.4L16 8l-6-5.9z"/>`),
		g.Group(children),
	)
}