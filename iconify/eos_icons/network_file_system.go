package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkFileSystem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 17h-3v-4h-2v4H8v1H2v2h6v1h8v-1h6v-2h-6v-1zm-1-4h3a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1h-6l-1-1H6a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h9Z"/>`),
		g.Group(children),
	)
}