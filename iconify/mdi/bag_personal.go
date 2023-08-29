package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BagPersonal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 5V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v1a4 4 0 0 0-4 4v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9a4 4 0 0 0-4-4m-6-1h4v1h-4V4m2 5l2 2l-2 2l-2-2l2-2m6 7H9v2H8v-2H6v-1h12v1Z"/>`),
		g.Group(children),
	)
}