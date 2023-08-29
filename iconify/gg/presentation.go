package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Presentation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm8-10h-2a6 6 0 0 0-12 0H4a8 8 0 1 1 16 0ZM4.252 14h15.496a8.003 8.003 0 0 1-15.496 0ZM8 12a4 4 0 1 1 8 0H8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}