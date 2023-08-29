package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravestiLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.952 8.537A7.5 7.5 0 1 1 7.538 9.95L4.663 7.075L2.188 9.55L.773 8.136l6.364-6.364l1.415 1.414l-2.475 2.475l2.875 2.876ZM13.502 20a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Z"/>`),
		g.Group(children),
	)
}