package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M18 17c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2H6c-1.103 0-2 .897-2 2v10c0 1.103.897 2 2 2h12zM4 19h16v2H4z" fill="currentColor"/>`),
		g.Group(children),
	)
}