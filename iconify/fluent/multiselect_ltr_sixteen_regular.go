package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiselectLtrSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.854 2.146a.5.5 0 0 1 0 .708l-2 2a.5.5 0 0 1-.708 0l-1-1a.5.5 0 1 1 .708-.708l.646.647l1.646-1.647a.5.5 0 0 1 .708 0ZM14.5 4h-8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1ZM6 8a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 0 1h-8A.5.5 0 0 1 6 8Zm-1.146 3.146a.5.5 0 0 1 0 .708l-2 2a.5.5 0 0 1-.708 0l-1-1a.5.5 0 0 1 .708-.708l.646.647l1.646-1.647a.5.5 0 0 1 .708 0ZM14.5 13h-8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1Z"/>`),
		g.Group(children),
	)
}