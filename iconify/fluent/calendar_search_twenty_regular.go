package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSearchTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 5.5A2.5 2.5 0 0 0 14.5 3h-9A2.5 2.5 0 0 0 3 5.5v3.756c.318-.112.653-.19 1-.229V7h12v7.5a1.5 1.5 0 0 1-1.5 1.5H9.121l1 1H14.5a2.5 2.5 0 0 0 2.5-2.5v-9ZM5.5 4h9A1.5 1.5 0 0 1 16 5.5V6H4v-.5A1.5 1.5 0 0 1 5.5 4Zm1.096 12.303a3.5 3.5 0 1 1 .707-.707l2.55 2.55a.5.5 0 0 1-.707.708l-2.55-2.55ZM7 13.5a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0Z"/>`),
		g.Group(children),
	)
}