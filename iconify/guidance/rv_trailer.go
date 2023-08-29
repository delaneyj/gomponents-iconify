package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RvTrailer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M19.5 7.5v.429c0 3.335 1.288 4.215 4 5.571v5H21m-1.5-11h2V7s-1.5-1-2-2.5H.5v14h2m17-11h-5v4h6.019M16 18.5a2.5 2.5 0 0 0 5 0m-5 0a2.5 2.5 0 0 1 5 0m-5 0H7.5m-5 0a2.5 2.5 0 0 0 5 0m-5 0a2.5 2.5 0 0 1 5 0m4-7v-4h-8v4h8Z"/>`),
		g.Group(children),
	)
}