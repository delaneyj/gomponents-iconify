package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17V7h2v4h4V7h2v10H9v-4H5v4H3Zm15 0v-3h-5V7h2v5h3V7h2v5h2v2h-2v3h-2Z"/>`),
		g.Group(children),
	)
}