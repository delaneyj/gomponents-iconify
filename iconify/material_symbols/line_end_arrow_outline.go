package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineEndArrowOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19v-6H2v-2h9V5l11 7l-11 7Zm2-3.65L18.275 12L13 8.65v6.7ZM13 12Z"/>`),
		g.Group(children),
	)
}