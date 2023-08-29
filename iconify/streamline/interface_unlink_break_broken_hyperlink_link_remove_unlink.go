package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUnlinkBreakBrokenHyperlinkLinkRemoveUnlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.54 9.7l2-2l2.12-2.12a3 3 0 0 0 0-4.24h0a3 3 0 0 0-4.24 0L7 2.76L5.46 4.3M3.5 6.26L1.38 8.38a3 3 0 0 0 0 4.24h0a3 3 0 0 0 4.24 0l1-1M4.17 2.05l5.66 9.9"/>`),
		g.Group(children),
	)
}