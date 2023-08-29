package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 10a5.002 5.002 0 0 1-4 4.9V17h2v2H9v-2h2v-2.1A5.002 5.002 0 0 1 7 10V5h10v5Zm-2-3H9v3a3 3 0 1 0 6 0V7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}