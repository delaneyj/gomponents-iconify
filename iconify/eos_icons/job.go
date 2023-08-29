package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Job(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 16.28V11l3-3l-4-4l-3 3H8V6h3V2H3v4h3v6H3v4h3v5h4.28A2 2 0 0 0 12 22a2 2 0 1 0-1.72-3H8v-3h3v-4H8V9h7l2 2v5.27a2 2 0 1 0 2 0ZM4 5V3h6v2Zm8 14a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm-2-6v2H4v-2Zm8-7.59L20.59 8L18 10.59L15.41 8ZM18 19a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}