package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightUpFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.05 12.361l-5.656 5.657l-1.415-1.414l5.657-5.657l-4.95-4.95H18v11.314l-4.95-4.95Z"/>`),
		g.Group(children),
	)
}