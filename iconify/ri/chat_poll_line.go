package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatPollLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H6.455L2 22.5V4a1 1 0 0 1 1-1h18Zm-1 2H4v13.385L5.763 17H20V5Zm-7 2v8h-2V7h2Zm4 2v6h-2V9h2Zm-8 2v4H7v-4h2Z"/>`),
		g.Group(children),
	)
}