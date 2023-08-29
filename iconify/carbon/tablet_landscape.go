package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletLandscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 13h2v6h-2z"/><path fill="currentColor" d="M30 7v18a2.002 2.002 0 0 1-2 2H4a2.002 2.002 0 0 1-2-2V7a2.002 2.002 0 0 1 2-2h24a2.003 2.003 0 0 1 2 2ZM4 25h24V7H4Z"/>`),
		g.Group(children),
	)
}