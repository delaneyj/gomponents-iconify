package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneFindLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 2a1 1 0 0 1 1 1v8h-2V4H7v16h4v2H6a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h12Zm-3 10a4 4 0 0 1 3.446 6.032l2.21 2.21l-1.413 1.415l-2.212-2.21A4 4 0 1 1 15 12Zm0 2a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		g.Group(children),
	)
}