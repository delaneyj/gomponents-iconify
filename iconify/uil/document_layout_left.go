package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentLayoutLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 12h6a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1Zm1-6h4v4H4Zm9 2h8a1 1 0 0 0 0-2h-8a1 1 0 0 0 0 2Zm0 10H3a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2Zm8-4H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2Zm0-4h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}