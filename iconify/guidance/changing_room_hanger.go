package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChangingRoomHanger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M9.5 6V5a2.5 2.5 0 0 1 5 0v.216a3.1 3.1 0 0 1-.908 2.192a26 26 0 0 1-12.079 6.839L.5 14.5v3l.782.26a33.892 33.892 0 0 0 21.436 0l.782-.26v-3l-1.013-.253A26 26 0 0 1 14 10.405"/>`),
		g.Group(children),
	)
}