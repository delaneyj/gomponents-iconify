package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongArrowReturn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M363 85h42v128H82l76 77l-30 30L0 192L128 64l30 30l-76 77h281V85z"/>`),
		g.Group(children),
	)
}