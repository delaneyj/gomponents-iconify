package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m20.1 8.1l4.243 4.243M30 4v6v-6Zm9.9 4.1l-4.243 4.243L39.9 8.101ZM44 18h-6h6Zm-4.1 9.9l-4.243-4.243l4.243 4.242ZM30 32v-6v6Zm-9.9-4.1l4.243-4.243l-4.242 4.242ZM16 18h6h-6Zm13.586.414L5.544 42.456"/>`),
		g.Group(children),
	)
}