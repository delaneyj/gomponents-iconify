package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213 21h128v128l-49-49l-61 62l-30-30l61-62zm-85 0L79 70l113 113v180h-43V201L49 100L0 149V21h128z"/>`),
		g.Group(children),
	)
}