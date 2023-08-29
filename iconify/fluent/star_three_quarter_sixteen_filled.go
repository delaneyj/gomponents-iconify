package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarThreeQuarterSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.194 2.101a.9.9 0 0 1 1.614 0L10 4.517V13.4l-1.999-1.051l-3.042 1.599a.9.9 0 0 1-1.306-.949l.58-3.387l-2.46-2.4a.9.9 0 0 1 .499-1.534l3.4-.495l1.522-3.082Z"/>`),
		g.Group(children),
	)
}