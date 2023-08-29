package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefundFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.005 7h-20V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v3Zm0 2v11a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1V9h20Zm-11 5v-2.5l-4.5 4.5h10.5v-2h-6Z"/>`),
		g.Group(children),
	)
}