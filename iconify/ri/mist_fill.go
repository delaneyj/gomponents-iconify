package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MistFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h4v2H4V4Zm12 15h4v2h-4v-2ZM2 9h10v2H2V9Zm12 0h6v2h-6V9ZM4 14h6v2H4v-2Zm8 0h10v2H12v-2ZM10 4h12v2H10V4ZM2 19h12v2H2v-2Z"/>`),
		g.Group(children),
	)
}