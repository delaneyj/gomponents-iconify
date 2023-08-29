package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundDetectionGlassBreakOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3ZM5 5v9l4.5-4.525l4 4L19 7.3V5H5Zm0 14h14v-8.7l-5.5 6.175l-4-4L5 17v2Z"/>`),
		g.Group(children),
	)
}