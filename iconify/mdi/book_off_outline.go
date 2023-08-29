package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 14.8l-2 2V4a2 2 0 0 1 2-2h12c.24 0 .47.04.68.12L16.8 4H13v3.8L10.79 10l-.29-.25L8 12V4H6v10.8M22.89 3L20 5.89V20c0 1.11-.89 2-2 2H6c-.58 0-1.1-.25-1.46-.65l-1.38 1.38l-1.27-1.27L21.61 1.73L22.89 3M18 7.89l-12 12V20h12V7.89Z"/>`),
		g.Group(children),
	)
}