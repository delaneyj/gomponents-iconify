package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoInputHdmi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 7V4a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v3H5v6l3 6v3h8v-3l3-6V7h-1M8 4h8v3h-2V5h-1v2h-2V5h-1v2H8V4Z"/>`),
		g.Group(children),
	)
}