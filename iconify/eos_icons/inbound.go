package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8v-2H4V6l8 5l8-5v8h2V4a2 2 0 0 0-2-2Zm-8 7L4 4h16Zm6 4v3h4v2h-4v3l-4-4Z"/>`),
		g.Group(children),
	)
}