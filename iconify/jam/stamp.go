package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 18v-2H2v2h16zM7 14V9a5 5 0 1 1 6 0v5h5a2 2 0 0 1 2 2v4H0v-4a2 2 0 0 1 2-2h5zm4 0V8.001l.799-.6a3 3 0 1 0-3.598 0l.799.6V14h2z"/>`),
		g.Group(children),
	)
}