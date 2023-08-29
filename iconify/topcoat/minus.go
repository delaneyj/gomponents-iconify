package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M39.5 21.5v-3c0-1.48-.43-2-2-2h-34c-1.48 0-2 .49-2 2v3c0 1.55.52 2 2 2h34c1.51 0 2-.48 2-2z"/>`),
		g.Group(children),
	)
}