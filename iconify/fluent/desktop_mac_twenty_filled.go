package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopMacTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 2a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h4v1a1 1 0 0 1-1 1h-.5a.5.5 0 0 0 0 1h7a.5.5 0 0 0 0-1H13a1 1 0 0 1-1-1v-1h4a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H4ZM3 13v-1h14v1a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1Zm5.732 4A1.99 1.99 0 0 0 9 16v-1h2v1c0 .364.097.706.268 1H8.732Z"/>`),
		g.Group(children),
	)
}