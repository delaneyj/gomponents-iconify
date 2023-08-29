package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M196.24 68.24L78.48 186H168a6 6 0 0 1 0 12H64a6 6 0 0 1-6-6V88a6 6 0 0 1 12 0v89.52L187.76 59.76a6 6 0 0 1 8.48 8.48Z"/>`),
		g.Group(children),
	)
}