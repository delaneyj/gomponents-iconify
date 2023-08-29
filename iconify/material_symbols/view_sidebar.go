package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewSidebar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 8V4h4v4h-4Zm0 6v-4h4v4h-4ZM2 20V4h14v16H2Zm16 0v-4h4v4h-4Z"/>`),
		g.Group(children),
	)
}