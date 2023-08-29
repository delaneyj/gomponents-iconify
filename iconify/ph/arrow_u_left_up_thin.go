package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowULeftUpThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 80v88a60 60 0 0 1-120 0V41.66L42.83 82.83a4 4 0 0 1-5.66-5.66l48-48a4 4 0 0 1 5.66 0l48 48a4 4 0 0 1-5.66 5.66L92 41.66V168a52 52 0 0 0 104 0V80a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}