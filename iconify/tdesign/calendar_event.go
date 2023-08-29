package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarEvent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 1v3h8V1h2v3h4v18H2V4h4V1h2ZM4 6v3h16V6H4Zm16 5H4v9h16v-9Zm-7.5 1.5h6v6h-6v-6Zm2 2v2h2v-2h-2Z"/>`),
		g.Group(children),
	)
}