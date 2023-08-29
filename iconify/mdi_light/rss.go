package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 16a3 3 0 0 1 3 3a3 3 0 0 1-3 3a3 3 0 0 1-3-3a3 3 0 0 1 3-3m0 1a2 2 0 0 0-2 2a2 2 0 0 0 2 2a2 2 0 0 0 2-2a2 2 0 0 0-2-2m-3-6c6.08 0 11 4.92 11 11h-1A10 10 0 0 0 2 12v-1m0-4a15 15 0 0 1 15 15h-1A14 14 0 0 0 2 8V7m0-4c10.5 0 19 8.5 19 19h-1A18 18 0 0 0 2 4V3Z"/>`),
		g.Group(children),
	)
}