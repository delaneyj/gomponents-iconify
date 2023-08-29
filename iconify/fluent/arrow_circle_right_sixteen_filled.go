package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleRightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14Zm3.462-6.809a.5.5 0 0 1-.106.16l-.003.003l-2.5 2.5a.5.5 0 0 1-.707-.708L9.793 8.5H5a.5.5 0 0 1 0-1h4.793L8.146 5.854a.5.5 0 1 1 .708-.708l2.5 2.5l.002.003a.499.499 0 0 1 .106.542Z"/>`),
		g.Group(children),
	)
}