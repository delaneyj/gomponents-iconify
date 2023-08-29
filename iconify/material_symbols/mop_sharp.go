package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MopSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.55 21H6v-3h2v3h3v-3h2v3h3v-3h2v3h2.45l-1-4H4.55l-1 4ZM1 23l2-8v-4h6V4q0-1.25.875-2.125T12 1q1.25 0 2.125.875T15 4v7h6v4l2 8H1Z"/>`),
		g.Group(children),
	)
}