package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLinkCreateHyperlinkLinkMakeUnlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 6.49L.79 9.67a1 1 0 0 0 0 1.42l2.12 2.12a1 1 0 0 0 1.42 0L7.51 10M10 7.51l3.18-3.18a1 1 0 0 0 0-1.42L11.09.79a1 1 0 0 0-1.42 0L6.49 4M9 5L5 9"/>`),
		g.Group(children),
	)
}