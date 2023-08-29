package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.7 14.01c-.23 0-.45-.06-.66-.18l-3.29-1.94a.515.515 0 0 0-.51 0l-3.29 1.94c-.41.24-.89.24-1.3.01c-.41-.23-.65-.66-.65-1.13V3.49c0-.83.67-1.5 1.5-1.5h8c.83 0 1.5.67 1.5 1.5v9.22a1.293 1.293 0 0 1-1.29 1.3Zm-4.2-3.2c.26 0 .53.07.76.21l3.29 1.94c.13.08.25.03.3 0c.05-.03.15-.1.15-.26V3.49c0-.28-.22-.5-.5-.5h-8c-.28 0-.5.22-.5.5v9.22c0 .16.1.23.15.26c.05.03.16.08.3 0l3.29-1.94c.24-.14.5-.21.76-.21Z"/>`),
		g.Group(children),
	)
}