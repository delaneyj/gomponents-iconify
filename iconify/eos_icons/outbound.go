package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outbound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8v-2H4V6l8 5l8-5v7h2V4a2 2 0 0 0-2-2Zm-8 7L4 4h16Zm10 8l-4 4v-3h-4v-2h4v-3Z"/>`),
		g.Group(children),
	)
}