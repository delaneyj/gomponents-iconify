package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16zm-1.828-8l3.535 3.536a1 1 0 0 1-1.414 1.414L6.05 10.707a1 1 0 0 1 0-1.414l4.243-4.243a1 1 0 0 1 1.414 1.414L8.172 10z"/>`),
		g.Group(children),
	)
}