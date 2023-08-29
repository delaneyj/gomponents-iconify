package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17h18v2H3v-2Zm0-6h3v3H3v-3Zm5 0h3v3H8v-3ZM3 5h3v3H3V5Zm10 0h3v3h-3V5Zm5 0h3v3h-3V5Zm-5 6h3v3h-3v-3Zm5 0h3v3h-3v-3ZM8 5h3v3H8V5Z"/>`),
		g.Group(children),
	)
}