package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Caravan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 7a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h2a3 3 0 0 0 3 3a3 3 0 0 0 3-3h8v-2h-2V9a2 2 0 0 0-2-2H5m0 2h5v3H5V9m8 0h4v3h-4V9m-3 7a1 1 0 0 1 1 1a1 1 0 0 1-1 1a1 1 0 0 1-1-1a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}