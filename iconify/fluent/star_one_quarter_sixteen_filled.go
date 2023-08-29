package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOneQuarterSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m6 4.52l-.327.663l-3.401.495a.898.898 0 0 0-.5 1.535l2.462 2.399l-.581 3.387a.899.899 0 0 0 1.306.949L6 13.401V4.52Z"/>`),
		g.Group(children),
	)
}