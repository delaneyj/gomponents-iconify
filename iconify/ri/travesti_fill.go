package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravestiFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.538 9.95L4.663 7.076L2.188 9.55L.773 8.136l6.364-6.364l1.415 1.414l-2.475 2.475l2.875 2.876A7.5 7.5 0 1 1 7.538 9.95Z"/>`),
		g.Group(children),
	)
}