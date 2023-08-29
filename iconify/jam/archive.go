package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5h16V2H2v3zm16 2v12a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V7H1a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1h-1zM4 18h12V7H4v11zm4-4h4a1 1 0 0 1 0 2H8a1 1 0 0 1 0-2z"/>`),
		g.Group(children),
	)
}