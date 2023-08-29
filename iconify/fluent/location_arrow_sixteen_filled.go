package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationArrowSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.637 2.07c.81-.312 1.604.483 1.293 1.292l-3.846 9.998c-.348.906-1.653.834-1.9-.105l-1.06-4.024a.5.5 0 0 0-.356-.357L2.744 7.815c-.94-.247-1.01-1.552-.105-1.9l9.998-3.846Z"/>`),
		g.Group(children),
	)
}