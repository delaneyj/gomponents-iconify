package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTrendingFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 6A1.5 1.5 0 0 1 9 7.5v28a3.5 3.5 0 0 0 3.5 3.5h28a1.5 1.5 0 0 1 0 3h-28A6.5 6.5 0 0 1 6 35.5v-28A1.5 1.5 0 0 1 7.5 6ZM38 17.121V23.5a1.5 1.5 0 0 0 3 0v-10a1.5 1.5 0 0 0-1.5-1.5h-10a1.5 1.5 0 0 0 0 3h6.379l-8.428 8.428l-4.442-4.038a1.5 1.5 0 0 0-2.037.018l-8.5 8a1.5 1.5 0 1 0 2.056 2.184l7.49-7.049l4.473 4.067a1.5 1.5 0 0 0 2.07-.05L38 17.122Z"/>`),
		g.Group(children),
	)
}