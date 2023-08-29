package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M23.712 15.703a1 1 0 0 0-1.423-1.406l-7.288 7.376V3.998a1 1 0 1 0-2 0v17.673l-7.287-7.374a1 1 0 0 0-1.422 1.406l8.82 8.927a1.25 1.25 0 0 0 1.779 0l8.821-8.927Z"/>`),
		g.Group(children),
	)
}