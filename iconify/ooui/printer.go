package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 1h10v4H5zM3 6a2 2 0 0 0-2 2v7h4v4h10v-4h4V8a2 2 0 0 0-2-2zm11 12H6v-6h8zm2-8a1 1 0 1 1 1-1a1 1 0 0 1-1 1z"/>`),
		g.Group(children),
	)
}