package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrangeSendBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 7h4v1H7.71l6.71 6.72l-.7.7L7 8.71V11H6V7m14 14H8v-9l1 1v7h10V10h-7l-1-1h9v12M3 16V4h12v3h-1V5H4v10h2v1H3Z"/>`),
		g.Group(children),
	)
}