package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundDetectionGlassBreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16V3h18v2.05l-7.5 8.425l-4-4L3 16Zm0 5v-2l6.5-6.525l4 4L21 8.05V21H3Z"/>`),
		g.Group(children),
	)
}