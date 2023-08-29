package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleUpF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 9.172l3.536 3.535a1 1 0 0 0 1.414-1.414L10.707 7.05a1 1 0 0 0-1.414 0L5.05 11.293a1 1 0 0 0 1.414 1.414L10 9.172zM10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10z"/>`),
		g.Group(children),
	)
}