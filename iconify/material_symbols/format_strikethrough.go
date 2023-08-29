package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatStrikethrough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 14v-2h20v2H2Zm8.5-4V7H5V4h14v3h-5.5v3h-3Zm0 10v-4h3v4h-3Z"/>`),
		g.Group(children),
	)
}