package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsExpandVerticalDiamondArrowDiamondDataVerticalDataInternetTransferNetwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="9.82" height="9.82" x="2.09" y="2.09" rx="1.07" transform="rotate(-45 7 7)"/><path d="M5.5 5.5L7 4l1.5 1.5m-3 3L7 10l1.5-1.5"/></g>`),
		g.Group(children),
	)
}