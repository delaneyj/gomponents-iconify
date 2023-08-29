package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WasteBasketEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9 4l-1.17 7H3.17L2 4h7zm.25-1.75A.25.25 0 0 1 9 2.5H2A.25.25 0 1 1 2 2h2V0h3v2h2a.25.25 0 0 1 .25.25zM6.5 2V.5h-2V2h2z" fill="currentColor"/>`),
		g.Group(children),
	)
}