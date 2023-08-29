package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneChatTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 4.25A2.25 2.25 0 0 0 15.75 2h-7.5A2.25 2.25 0 0 0 6 4.25v15.5A2.25 2.25 0 0 0 8.25 22h2.837l.456-1.5H8.25a.75.75 0 0 1-.75-.75V4.25a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 .75.75v6.827a6.551 6.551 0 0 1 1.5-.058V4.25Zm3 9.006A5.475 5.475 0 0 0 17.501 12a5.501 5.501 0 0 0-4.812 8.169l-.666 2.186a.5.5 0 0 0 .624.625l2.187-.666A5.501 5.501 0 0 0 21 13.256ZM15.5 17a.5.5 0 1 1 0-1h4.002a.5.5 0 1 1 0 1H15.5Zm2.001 2h-2a.5.5 0 1 1 0-1h2a.5.5 0 1 1 0 1Z"/>`),
		g.Group(children),
	)
}