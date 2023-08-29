package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 26.59L21.41 24L20 25.41l4 4l7-7L29.59 21L24 26.59zM15 23h2v2h-2zm-6 0h2v2H9zm12-5h2v2h-2zm-6 0h2v2h-2zm-6 0h2v2H9zm12-5h2v2h-2zm-6 0h2v2h-2zm-6 0h2v2H9zm0-6h14v3H9z"/><path fill="currentColor" d="M17 30H6.005A2.007 2.007 0 0 1 4 27.995V3.996A1.998 1.998 0 0 1 5.996 2h20.008A1.998 1.998 0 0 1 28 3.996V18h-2V4H6v24h11Z"/>`),
		g.Group(children),
	)
}