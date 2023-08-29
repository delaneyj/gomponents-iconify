package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApprovalsAppTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.543 1.793a1 1 0 0 1 1.414 0l2.5 2.5a1 1 0 0 1 0 1.414l-2.5 2.5a1 1 0 1 1-1.414-1.414l.758-.759a7 7 0 1 0 7.645 7.842a1 1 0 1 1 1.984.248a9 9 0 1 1-9.572-10.101l-.815-.816a1 1 0 0 1 0-1.414Zm5.664 8a1 1 0 0 1 0 1.414l-4.5 4.5a1 1 0 0 1-1.414 0l-2-2a1 1 0 1 1 1.414-1.414L11 13.586l3.793-3.793a1 1 0 0 1 1.414 0Z"/>`),
		g.Group(children),
	)
}