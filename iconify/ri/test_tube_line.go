package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTubeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 2v2h-1v14a4 4 0 0 1-8 0V4H7V2h10Zm-3 8h-4v8a2 2 0 1 0 4 0v-8Zm-1 5a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm-2-3a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm3-8h-4v4h4V4Z"/>`),
		g.Group(children),
	)
}