package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2v11h3v9l7-12h-4l4-8m2 13h2v2h-2v-2m0-8h2v6h-2V7Z"/>`),
		g.Group(children),
	)
}