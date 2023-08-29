package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsMoveHorizontalCircleTransferDataInternetArrowHorizontalNetworkCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m6 10.5l-2-2h6m-2-5l2 2H4"/><circle cx="7" cy="7" r="6.5"/></g>`),
		g.Group(children),
	)
}