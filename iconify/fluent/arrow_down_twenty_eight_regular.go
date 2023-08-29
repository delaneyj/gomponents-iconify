package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M23.791 15.267a.75.75 0 1 0-1.08-1.04L14.75 22.5V3.748a.75.75 0 0 0-1.5 0V22.5l-7.959-8.273a.75.75 0 0 0-1.08 1.04l9.069 9.428a1 1 0 0 0 1.441 0l9.07-9.428Z"/>`),
		g.Group(children),
	)
}