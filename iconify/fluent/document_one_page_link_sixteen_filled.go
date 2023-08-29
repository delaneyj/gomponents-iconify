package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentOnePageLinkSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 3a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v6.5a1.13 1.13 0 0 0-.083.003h-2.834A1.013 1.013 0 0 0 10 9.5h-.5A3 3 0 0 0 7.841 15H5a2 2 0 0 1-2-2V3Zm6.59 1.992A.5.5 0 0 0 9.5 4h-3l-.09.008A.5.5 0 0 0 6.5 5h3l.09-.008ZM10 8a.5.5 0 0 0-.5-.5h-3l-.09.008a.5.5 0 0 0 .09.992h3l.09-.008A.5.5 0 0 0 10 8Zm-3 4.5A2.5 2.5 0 0 1 9.5 10h.5a.5.5 0 0 1 0 1h-.5a1.5 1.5 0 0 0 0 3h.5a.5.5 0 0 1 0 1h-.5A2.5 2.5 0 0 1 7 12.5Zm5.5-2a.5.5 0 0 1 .5-.5h.5a2.5 2.5 0 0 1 0 5H13a.5.5 0 0 1 0-1h.5a1.5 1.5 0 0 0 0-3H13a.5.5 0 0 1-.5-.5Zm-3.5 2a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}