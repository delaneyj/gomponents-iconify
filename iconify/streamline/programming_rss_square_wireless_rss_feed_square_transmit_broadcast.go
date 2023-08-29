package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingRssSquareWirelessRssFeedSquareTransmitBroadcast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><circle cx="4" cy="10" r=".5"/><path d="M4.5 6.5a3 3 0 0 1 3 3m-3-6a6 6 0 0 1 6 6"/></g>`),
		g.Group(children),
	)
}