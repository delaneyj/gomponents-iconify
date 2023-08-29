package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17V7h2v4h4V7h2v10h-2v-4H7v4H5Zm12 0V9h-2V7h4v10h-2Z"/>`),
		g.Group(children),
	)
}