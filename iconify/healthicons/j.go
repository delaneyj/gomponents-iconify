package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func J(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M30 10a2 2 0 0 1 2 2v18a8 8 0 1 1-16 0a2 2 0 1 1 4 0a4 4 0 1 0 8 0V12a2 2 0 0 1 2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}