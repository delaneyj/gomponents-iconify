package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalFourG(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 16.5v-6h-4.5v3H19v3h-3v-9h6v-3h-6a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h3a3 3 0 0 0 3-3m-14 3h3v-15H8v6H5v-6H2v9h6v6Z"/>`),
		g.Group(children),
	)
}