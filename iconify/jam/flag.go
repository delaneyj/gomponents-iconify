package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 1h12a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H2v7a1 1 0 0 1-2 0V1a1 1 0 1 1 2 0zm0 9h12V3H2v7z"/>`),
		g.Group(children),
	)
}