package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrayMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 10H8V8h8M2 17a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-5h-2v5H4v-5H2Z"/>`),
		g.Group(children),
	)
}