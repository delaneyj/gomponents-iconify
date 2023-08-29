package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentDataLinkTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 8.5a.5.5 0 0 0-1 0V13h.5c.17 0 .337.012.5.035V8.5ZM5 4v9H4V4a2 2 0 0 1 2-2h4.586a1.5 1.5 0 0 1 1.06.44l3.915 3.914A1.5 1.5 0 0 1 16 7.414V16a2 2 0 0 1-2 2h-3.337c.148-.31.251-.647.302-1H14a1 1 0 0 0 1-1V8h-3.5A1.5 1.5 0 0 1 10 6.5V3H6a1 1 0 0 0-1 1Zm5.5 8.5v2.196a3.519 3.519 0 0 0-1-1.069V12.5a.5.5 0 0 1 1 0Zm2.5-2a.5.5 0 0 0-1 0v5a.5.5 0 0 0 1 0v-5ZM11.5 7h3.293L11 3.207V6.5a.5.5 0 0 0 .5.5Zm-8 7a2.5 2.5 0 0 0 0 5H4a.5.5 0 0 0 0-1h-.5a1.5 1.5 0 0 1 0-3H4a.5.5 0 0 0 0-1h-.5ZM7 14a.5.5 0 0 0 0 1h.5a1.5 1.5 0 0 1 0 3H7a.5.5 0 0 0 0 1h.5a2.5 2.5 0 0 0 0-5H7Zm-4 2.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}