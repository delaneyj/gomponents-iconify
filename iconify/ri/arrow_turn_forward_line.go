package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnForwardLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 5.828V13a6 6 0 0 1-12 0V4H3v9a8 8 0 1 0 16 0V5.828l2.536 2.536L22.95 6.95L18 2l-4.95 4.95l1.415 1.414L17 5.828Z"/>`),
		g.Group(children),
	)
}