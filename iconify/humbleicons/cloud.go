package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9.036a3.485 3.485 0 0 1 1.975.99M4 12.5A3.5 3.5 0 0 0 7.5 16h9.75a2.75 2.75 0 0 0 .734-5.4A5 5 0 0 0 8.37 9.108A3.5 3.5 0 0 0 4 12.5z"/>`),
		g.Group(children),
	)
}