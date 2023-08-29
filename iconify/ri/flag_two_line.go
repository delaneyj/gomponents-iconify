package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.138 3a.5.5 0 0 1 .434.748L18 10l3.573 6.252a.5.5 0 0 1-.435.748H4v5H2V3h19.138Zm-2.584 2H4v10h14.554l-2.857-5l2.857-5Z"/>`),
		g.Group(children),
	)
}