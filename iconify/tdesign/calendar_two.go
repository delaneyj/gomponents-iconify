package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 1v3h8V1h2v3h4v18H2V4h4V1h2ZM4 6v3h16V6H4Zm16 5H4v9h16v-9Z"/><path fill="currentColor" d="m16.914 13.25l-5.657 5.657l-3.535-3.536l1.414-1.414l2.121 2.121l4.243-4.242l1.414 1.414Z"/>`),
		g.Group(children),
	)
}