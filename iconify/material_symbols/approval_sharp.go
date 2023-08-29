package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApprovalSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22v-8h16v8H4Zm2-4h12v-2H6v2Zm6-4L7 7q0-2.075 1.463-3.538T12 2q2.075 0 3.538 1.463T17 7l-5 7Z"/>`),
		g.Group(children),
	)
}