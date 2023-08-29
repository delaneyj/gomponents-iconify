package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.636 13.052L5.98 7.395L7.394 5.98l5.657 5.657L18 6.687v11.314H6.686l4.95-4.95Z"/>`),
		g.Group(children),
	)
}