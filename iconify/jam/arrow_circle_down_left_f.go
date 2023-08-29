package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDownLeftF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.172 12.828a.997.997 0 0 0 1 1h6a1 1 0 1 0 0-2H9.586l3.95-3.95a1 1 0 1 0-1.415-1.414l-3.95 3.95V6.828a1 1 0 1 0-2 0v6zM10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10z"/>`),
		g.Group(children),
	)
}