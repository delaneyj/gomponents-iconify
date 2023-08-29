package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 1H3a2 2 0 0 0-2 2v9h2V3h14Z"/><path fill="currentColor" d="M21 5H7a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h5.53v1.53H11V22h6v-1.48h-1.52V19H21a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2Zm0 12H7V7h14Z"/>`),
		g.Group(children),
	)
}