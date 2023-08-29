package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentSplitHintSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 4.5V1H4.5A1.5 1.5 0 0 0 3 2.5v5a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V6H9.5A1.5 1.5 0 0 1 8 4.5Zm1 0V1.25L12.75 5H9.5a.5.5 0 0 1-.5-.5Zm-5 5a.5.5 0 0 0-1 0v1a.5.5 0 0 0 1 0v-1Zm9 0a.5.5 0 0 0-1 0v1a.5.5 0 0 0 1 0v-1Zm-9 3a.5.5 0 0 0-1 0v.5a2 2 0 0 0 2 2h.5a.5.5 0 0 0 0-1H5a1 1 0 0 1-1-1v-.5Zm9 0a.5.5 0 0 0-1 0v.5a1 1 0 0 1-1 1h-.5a.5.5 0 0 0 0 1h.5a2 2 0 0 0 2-2v-.5ZM7.5 14a.5.5 0 0 0 0 1h1a.5.5 0 0 0 0-1h-1Z"/>`),
		g.Group(children),
	)
}