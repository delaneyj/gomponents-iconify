package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailCheckLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 14h-2V7.238l-7.928 7.1L4 7.216V19h10v2H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v10ZM4.511 5l7.55 6.662L19.502 5H4.511ZM19 22l-3.536-3.535l1.415-1.415L19 19.172l3.535-3.536l1.415 1.414L19 22Z"/>`),
		g.Group(children),
	)
}