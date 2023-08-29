package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLinkBrokenBreakBrokenHyperlinkLinkRemoveUnlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.14 7.13L1.27 9a2.65 2.65 0 0 0 0 3.74h0a2.65 2.65 0 0 0 3.74 0l1.11-1.11M9 9.5h1.86a2.64 2.64 0 0 0 2.64-2.64h0a2.64 2.64 0 0 0-2.64-2.64H8.22M7 .5l-.5 2m-6 1l2 1m.5-4l1 2"/>`),
		g.Group(children),
	)
}