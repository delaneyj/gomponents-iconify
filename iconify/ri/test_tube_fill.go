package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTubeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 2v2h-1v14a4 4 0 0 1-8 0V4H7V2h10Zm-4 13a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-2-3a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm3-8h-4v4h4V4Z"/>`),
		g.Group(children),
	)
}