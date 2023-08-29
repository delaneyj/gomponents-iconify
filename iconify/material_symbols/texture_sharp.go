package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextureSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.075 20.925v-1.4l16.45-16.45h1.425v1.4L4.475 20.925h-1.4ZM3 14.7v-2.8L11.9 3h2.8L3 14.7ZM3 7V3h4L3 7Zm14 14l4-4v4h-4Zm-7.7 0L21 9.3v2.8L12.1 21H9.3Z"/>`),
		g.Group(children),
	)
}