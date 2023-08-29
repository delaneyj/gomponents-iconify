package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowResetFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.634 6.634a1.25 1.25 0 0 0-1.768-1.768l-7.5 7.5a1.25 1.25 0 0 0 0 1.768l7.5 7.5a1.25 1.25 0 0 0 1.768-1.768L8.268 14.5H26.5c7.18 0 13 5.82 13 13s-5.82 13-13 13s-13-5.82-13-13a1.25 1.25 0 1 0-2.5 0C11 36.06 17.94 43 26.5 43C35.06 43 42 36.06 42 27.5C42 18.94 35.06 12 26.5 12H8.268l5.366-5.366Z"/>`),
		g.Group(children),
	)
}