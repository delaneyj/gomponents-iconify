package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsJoystickAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M16 6H8a6 6 0 0 0 0 12h8a6 6 0 0 0 0-12zm-5 7H9v2H7v-2H5v-2h2V9h2v2h2v2zm3.5 2a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3zm3-3a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3z" fill="currentColor"/>`),
		g.Group(children),
	)
}