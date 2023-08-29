package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 1h16v5h-2V3H6v18h12v-5h2v7H4V1Zm9 7h11v2H13V8Zm0 4h8v2h-8v-2Zm-2 5h2.004v2.004H11V17Z"/>`),
		g.Group(children),
	)
}