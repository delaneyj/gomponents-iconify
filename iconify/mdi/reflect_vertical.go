package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReflectVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 23H6c-.39 0-.74-.21-.91-.55a.998.998 0 0 1 .1-1.05l6.01-8.02c.38-.51 1.22-.51 1.6 0l6.01 8.02c.23.3.27.71.1 1.05c-.17.34-.52.55-.91.55m0-22c.39 0 .74.21.91.55c.17.34.13.75-.1 1.05l-6.01 8.02c-.38.51-1.22.51-1.6 0L5.19 2.6a.998.998 0 0 1-.1-1.05C5.26 1.21 5.61 1 6 1h12M8 3l4 5.35L16 3H8Z"/>`),
		g.Group(children),
	)
}