package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintBrushBroadBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 20H72a44.05 44.05 0 0 0-44 44v72a28 28 0 0 0 28 28h39.64L92 207v1a36 36 0 0 0 72 0v-1l-3.6-43H200a28 28 0 0 0 28-28V32a12 12 0 0 0-12-12ZM72 44h88v24a12 12 0 0 0 24 0V44h20v52H52V64a20 20 0 0 1 20-20Zm128 96h-44a20 20 0 0 0-19.85 22.4l3.84 46a12 12 0 0 1-24 0l3.84-46A20 20 0 0 0 100 140H56a4 4 0 0 1-4-4v-16h152v16a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}