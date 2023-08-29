package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintBrushBroadFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 24H72a40 40 0 0 0-40 40v72a24 24 0 0 0 24 24h48l-7.89 46.67A8.42 8.42 0 0 0 96 208a32 32 0 0 0 64 0a8.42 8.42 0 0 0-.11-1.33L152 160h48a24 24 0 0 0 24-24V32a8 8 0 0 0-8-8ZM72 40h104v40a8 8 0 0 0 16 0V40h16v72H48V64a24 24 0 0 1 24-24Z"/>`),
		g.Group(children),
	)
}