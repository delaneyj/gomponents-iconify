package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swiss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 3v18h18V3H3Zm11 4h-4v3H7v4h3v3h4v-3h3v-4h-3V7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}