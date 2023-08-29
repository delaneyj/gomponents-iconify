package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftUpFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.361 10.947l5.657 5.657l-1.414 1.414l-5.657-5.657l-4.95 4.95V5.997h11.314l-4.95 4.95Z"/>`),
		g.Group(children),
	)
}