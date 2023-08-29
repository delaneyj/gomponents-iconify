package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphaB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 10.5V9a2 2 0 0 0-2-2H9v10h4a2 2 0 0 0 2-2v-1.5c0-.8-.7-1.5-1.5-1.5c.8 0 1.5-.7 1.5-1.5M13 15h-2v-2h2v2m0-4h-2V9h2v2Z"/>`),
		g.Group(children),
	)
}