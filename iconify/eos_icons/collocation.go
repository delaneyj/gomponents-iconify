package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Collocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm0 18H4V4h16Z"/><path fill="currentColor" d="M7.51 18A2.5 2.5 0 0 0 10 16h4.1a2.5 2.5 0 1 0 2.45-3h-.05l-2.42-4.1a2.5 2.5 0 1 0-4.14 0L7.56 13a2.5 2.5 0 0 0 0 5Zm3.17-8.39a2.46 2.46 0 0 0 2.66 0l2.1 3.64A2.49 2.49 0 0 0 14.06 15H10a2.49 2.49 0 0 0-1.38-1.75Z"/>`),
		g.Group(children),
	)
}