package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.211 12.733a.75.75 0 1 0 1.081 1.04l7.96-8.275v18.753a.75.75 0 1 0 1.5 0V5.5l7.958 8.274a.75.75 0 0 0 1.081-1.04l-9.069-9.428a1 1 0 0 0-1.441 0l-9.07 9.428Z"/>`),
		g.Group(children),
	)
}