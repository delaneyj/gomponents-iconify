package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaletteRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 15v2a3 3 0 0 0 3 3A10 10 0 1 0 0 10a5 5 0 0 0 5 5ZM3 8.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 1 1-3 0m3-4a1.5 1.5 0 1 1 3 0a1.5 1.5 0 1 1-3 0m5 0a1.5 1.5 0 1 1 3 0a1.5 1.5 0 1 1-3 0m3 4a1.5 1.5 0 1 1 3 0a1.5 1.5 0 1 1-3 0"/>`),
		g.Group(children),
	)
}