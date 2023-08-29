package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M225.86 110.48L57.8 14.58a20 20 0 0 0-28.64 24.09l30.61 89.21l-30.61 89.45A20 20 0 0 0 48 244a20.1 20.1 0 0 0 9.81-2.58l.09-.06l168-96.07a20 20 0 0 0 0-34.81ZM55.24 215.23L81 140h55a12 12 0 0 0 0-24H81.07L55.25 40.76l152.69 87.13Z"/>`),
		g.Group(children),
	)
}