package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteControlTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 2a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h12Zm-1 2H7v16h10V4Zm-2 11v2h-2v-2h2Zm-4 0v2H9v-2h2Zm2-9v2h2v2h-2.001L13 12h-2l-.001-2H9V8h2V6h2Z"/>`),
		g.Group(children),
	)
}