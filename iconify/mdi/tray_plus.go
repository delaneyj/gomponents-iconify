package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrayPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 12h2v5h16v-5h2v5a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2m9-12h2v3h3v2h-3v3h-2v-3H8V8h3Z"/>`),
		g.Group(children),
	)
}