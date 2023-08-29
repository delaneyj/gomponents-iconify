package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0h32v24H0zm2 2v20h28V2zm2.012 18.054S6.092 6.008 9.181 6.008s4.442 9.908 6.926 9.908s2.298-3.871 4.116-3.871s8.08 8.01 8.08 8.01zM25 10a3 3 0 1 1 0-6a3 3 0 0 1 0 6z"/>`),
		g.Group(children),
	)
}