package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Replica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 11h-6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2Zm0 8h-6v-6h6ZM12 3H6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h3v-2H6V5h6v4h2V5a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}