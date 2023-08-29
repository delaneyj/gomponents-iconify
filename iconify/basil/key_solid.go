package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 12a5.75 5.75 0 0 1 10.8-2.75H21c.966 0 1.75.784 1.75 1.75v2.5a.75.75 0 0 1-.75.75h-2.25V16a.75.75 0 0 1-.75.75h-2.5a.75.75 0 0 1-.75-.75v-1.75h-3.457A5.751 5.751 0 0 1 1.25 12ZM7 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}