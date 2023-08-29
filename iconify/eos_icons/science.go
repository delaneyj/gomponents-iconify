package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Science(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.68 22H5.32a2.318 2.318 0 0 1-1.96-3.55L4.27 17L9 9.46V3H7V1h10v2h-2v6.46L19.73 17l.91 1.45A2.318 2.318 0 0 1 18.68 22Z"/>`),
		g.Group(children),
	)
}