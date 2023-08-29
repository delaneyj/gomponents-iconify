package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortDescendingThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M44 128a4 4 0 0 1 4-4h72a4 4 0 0 1 0 8H48a4 4 0 0 1-4-4Zm4-60h56a4 4 0 0 0 0-8H48a4 4 0 0 0 0 8Zm136 120H48a4 4 0 0 0 0 8h136a4 4 0 0 0 0-8Zm42.83-102.83l-40-40a4 4 0 0 0-5.66 0l-40 40a4 4 0 0 0 5.66 5.66L180 57.66V144a4 4 0 0 0 8 0V57.66l33.17 33.17a4 4 0 1 0 5.66-5.66Z"/>`),
		g.Group(children),
	)
}