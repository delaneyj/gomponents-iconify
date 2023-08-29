package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarChatTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 8.5v3.521a6.501 6.501 0 0 0-9.385 8.243L11.39 21H6.25A3.25 3.25 0 0 1 3 17.75V8.5h18ZM17.75 3A3.25 3.25 0 0 1 21 6.25V7H3v-.75A3.25 3.25 0 0 1 6.25 3h11.5ZM21 13.256A5.475 5.475 0 0 0 17.501 12a5.501 5.501 0 0 0-4.812 8.169l-.666 2.186a.5.5 0 0 0 .624.625l2.187-.666A5.501 5.501 0 0 0 21 13.256ZM15.5 17a.5.5 0 1 1 0-1h4.002a.5.5 0 1 1 0 1H15.5Zm2.001 2h-2a.5.5 0 1 1 0-1h2a.5.5 0 1 1 0 1Z"/>`),
		g.Group(children),
	)
}