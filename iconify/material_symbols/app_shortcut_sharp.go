package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppShortcutSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.15 13H12v3h-2v-5h7.15L15.6 9.4L17 8l4 4l-4 4l-1.4-1.4l1.55-1.6ZM5 23V1h14v6h-2V6H7v12h10v-1h2v6H5Z"/>`),
		g.Group(children),
	)
}