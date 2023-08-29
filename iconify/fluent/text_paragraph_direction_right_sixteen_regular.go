package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextParagraphDirectionRightSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 3a2.5 2.5 0 0 0 0 5h1V3H9Zm2 0v10.5a.5.5 0 0 1-1 0V9H9a3.5 3.5 0 1 1 0-7h4.5a.5.5 0 0 1 0 1H13v10.5a.5.5 0 0 1-1 0V3h-1ZM2.146 6.146a.5.5 0 0 1 .708 0l1.5 1.5a.5.5 0 0 1 0 .708l-1.5 1.5a.5.5 0 0 1-.708-.708L3.293 8L2.146 6.854a.5.5 0 0 1 0-.708Z"/>`),
		g.Group(children),
	)
}