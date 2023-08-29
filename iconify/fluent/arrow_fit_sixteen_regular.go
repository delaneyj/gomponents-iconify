package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFitSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.854 5.146a.5.5 0 0 1 0 .708L2.707 7H6.5a.5.5 0 0 1 0 1H2.707l1.147 1.146a.5.5 0 1 1-.708.708l-2-2a.5.5 0 0 1 0-.708l2-2a.5.5 0 0 1 .708 0Zm8.292 0a.5.5 0 0 1 .708 0l2 2a.5.5 0 0 1 0 .708l-2 2a.5.5 0 0 1-.708-.708L13.293 8H9.5a.5.5 0 0 1 0-1h3.793l-1.147-1.146a.5.5 0 0 1 0-.708Z"/>`),
		g.Group(children),
	)
}