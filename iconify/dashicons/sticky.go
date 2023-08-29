package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sticky(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 3.61V1.04l8.99-.01l-.01 2.58a2.737 2.737 0 0 0-2.16 2.67v.5c.01 1.31.93 2.4 2.17 2.66l-.01 2.58h-3.41l-.01 2.57c0 .6-.47 4.41-1.06 4.41c-.6 0-1.08-3.81-1.08-4.41v-2.56L5 12.02l.01-2.58a2.707 2.707 0 0 0 2.15-2.66v-.5c0-1.31-.92-2.41-2.16-2.67z"/>`),
		g.Group(children),
	)
}