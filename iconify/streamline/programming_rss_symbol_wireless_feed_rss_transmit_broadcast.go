package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingRssSymbolWirelessFeedRssTransmitBroadcast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.13" cy="11.88" r="1.5"/><path d="M13.38 12.12A11.5 11.5 0 0 0 1.88.62m.24 4.76a6.5 6.5 0 0 1 6.5 6.5"/></g>`),
		g.Group(children),
	)
}