package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 10c0 3.9 3.1 7 7 7c1.4 0 2.7-.5 3.8-1.2L19 21l2-2l-5.2-5.2c.8-1.1 1.2-2.4 1.2-3.8c0-3.9-3.1-7-7-7s-7 3.1-7 7zm2 0c0-2.8 2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5s-5-2.2-5-5z"/><path fill="currentColor" d="M7 9h6v2H7z"/>`),
		g.Group(children),
	)
}