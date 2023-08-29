package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOneQuarterSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.808 2.101a.9.9 0 0 0-1.614 0L5.673 5.183l-3.401.495a.898.898 0 0 0-.5 1.535l2.462 2.399l-.581 3.387a.899.899 0 0 0 1.306.949L8 12.348l3.042 1.6a.9.9 0 0 0 1.306-.949l-.581-3.387l2.461-2.4a.9.9 0 0 0-.499-1.534l-3.4-.495l-1.522-3.082ZM6 12.271V6.136a.9.9 0 0 0 .546-.463l1.455-2.947l1.455 2.947a.9.9 0 0 0 .677.492l3.253.473l-2.354 2.294a.9.9 0 0 0-.258.797l.555 3.24l-2.91-1.53a.9.9 0 0 0-.837 0L6 12.271Z"/>`),
		g.Group(children),
	)
}