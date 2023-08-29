package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandContent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19v-6h2v4h4v2H5Zm12-8V7h-4V5h6v6h-2Z"/>`),
		g.Group(children),
	)
}