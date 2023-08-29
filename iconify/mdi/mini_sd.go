package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiniSd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-6l-2-2V6a2 2 0 0 0-2-2H6m1 2h2v4H7V6m3 0h2v4h-2V6m3 0h2v4h-2V6Z"/>`),
		g.Group(children),
	)
}