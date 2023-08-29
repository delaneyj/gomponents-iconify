package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUpRightF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.828 7.172a.997.997 0 0 0-1-1h-6a1 1 0 1 0 0 2h3.586l-3.95 3.95a1 1 0 0 0 1.415 1.414l3.95-3.95v3.586a1 1 0 0 0 2 0v-6zM10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10z"/>`),
		g.Group(children),
	)
}