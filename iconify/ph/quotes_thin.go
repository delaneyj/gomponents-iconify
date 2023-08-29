package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuotesThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M100 60H40a12 12 0 0 0-12 12v64a12 12 0 0 0 12 12h64v12a36 36 0 0 1-36 36a4 4 0 0 0 0 8a44.05 44.05 0 0 0 44-44V72a12 12 0 0 0-12-12Zm4 80H40a4 4 0 0 1-4-4V72a4 4 0 0 1 4-4h60a4 4 0 0 1 4 4Zm112-80h-60a12 12 0 0 0-12 12v64a12 12 0 0 0 12 12h64v12a36 36 0 0 1-36 36a4 4 0 0 0 0 8a44.05 44.05 0 0 0 44-44V72a12 12 0 0 0-12-12Zm4 80h-64a4 4 0 0 1-4-4V72a4 4 0 0 1 4-4h60a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}